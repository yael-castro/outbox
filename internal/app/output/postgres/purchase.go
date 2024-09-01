package postgres

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
)

func NewPurchaseSaver(db *sql.DB, errLogger *log.Logger) business.PurchaseSaver {
	return purchaseSaver{
		errLogger: errLogger,
		db:        db,
	}
}

type purchaseSaver struct {
	errLogger *log.Logger
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

	// Inserting message outbox
	_, err = tx.ExecContext(
		ctx,
		insertPurchaseOutbox,
		purchaseSQL.ID,
		purchaseSQL.OrderID,
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
