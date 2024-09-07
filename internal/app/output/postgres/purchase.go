//go:build http

package postgres

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
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
		purchase.OrderID,
	).Scan(&purchaseSQL.ID)
	if err != nil {
		p.errLogger.Printf("inserting purchase record: %v", err)
		return
	}

	message, err := purchaseSQL.MarshalBinary()
	if err != nil {
		p.errLogger.Printf("marshaling purchase record: %v", err)
		return
	}

	// Inserting outbox message
	_, err = tx.ExecContext(
		ctx,
		insertOutboxMessage,
		p.topic,
		nil,
		message,
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
