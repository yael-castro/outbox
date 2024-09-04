package postgres

import (
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
)

func NewPurchase(purchase *business.Purchase) Purchase {
	return Purchase{
		ID: sql.NullInt64{
			Int64: int64(purchase.ID),
			Valid: purchase.ID > 0,
		},
		OrderID: sql.NullInt64{
			Int64: int64(purchase.OrderID),
			Valid: purchase.OrderID > 0,
		},
	}
}

type Purchase struct {
	ID      sql.NullInt64
	OrderID sql.NullInt64
}

func (p *Purchase) ToBusiness() *business.Purchase {
	return &business.Purchase{
		ID:      uint64(p.ID.Int64),
		OrderID: uint64(p.OrderID.Int64),
	}
}

type PurchaseMessage struct {
	ID       sql.NullInt64
	Purchase Purchase
}

func (p PurchaseMessage) ToBusiness() business.PurchaseMessage {
	return business.PurchaseMessage{
		ID:       uint64(p.ID.Int64),
		Purchase: *p.Purchase.ToBusiness(),
	}
}
