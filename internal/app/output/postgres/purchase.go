//go:build http

package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
	"strconv"
)

type PurchaseSaverConfig struct {
	Topic  string
	DB     *sql.DB
	Logger *log.Logger
}

func NewPurchaseSaver(config PurchaseSaverConfig) business.PurchaseSaver {
	return purchaseSaver{
		errLogger: config.Logger,
		topic:     config.Topic,
		db:        config.DB,
	}
}

type purchaseSaver struct {
	errLogger *log.Logger
	topic     string
	db        *sql.DB
}

func (p purchaseSaver) SavePurchase(ctx context.Context, purchase *business.Purchase) (err error) {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		p.errLogger.Println("begin transaction")
		return
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// Building SQL record for Purchase entity
	purchaseSQL := NewPurchase(purchase)

	// Inserting purchase record
	err = tx.QueryRowContext(
		ctx,
		insertPurchase,
		purchaseSQL.OrderID,
	).Scan(&purchaseSQL.ID)
	if err != nil {
		p.errLogger.Printf("inserting purchase record: %[1]v (%[1]T)", err)

		// Error handling for postgres errors
		const violateUniqueConstraint = "23505"

		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code == violateUniqueConstraint {
			err = fmt.Errorf("%w: the order id already exists", business.ErrDuplicatedOrderID)
			return
		}

		return
	}

	message, err := purchaseSQL.MarshalBinary()
	if err != nil {
		p.errLogger.Printf("marshaling purchase record: %v", err)
		return
	}

	// Setting message headers
	headers, err := (&Headers{
		{
			Key:   "purchase_id",
			Value: []byte(strconv.FormatInt(purchaseSQL.OrderID.Int64, 10)),
		},
	}).MarshalBinary()
	if err != nil {
		return err
	}

	// Inserting outbox message
	_, err = tx.ExecContext(
		ctx,
		insertOutboxMessage,
		p.topic, // topic
		nil,     // partition_key
		headers, // headers
		message, // message_value
	)
	if err != nil {
		p.errLogger.Printf("inserting purchase message: %v", err)
		return
	}

	// Updating purchase id in the business entity
	*purchase = *purchaseSQL.ToBusiness()

	// Last step! Commit changes
	return tx.Commit()
}
