package postgres

import (
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
)

func NewPurchase(purchase *business.Purchase) Purchase {
	return Purchase{
		ID: sql.Null[int64]{
			V:     int64(purchase.ID),
			Valid: purchase.ID > 0,
		},
		OrderID: sql.Null[int64]{
			V:     int64(purchase.OrderID),
			Valid: purchase.OrderID > 0,
		},
	}
}

type Purchase struct {
	ID      sql.Null[int64]
	OrderID sql.Null[int64]
}

func (p *Purchase) ToBusiness() *business.Purchase {
	return &business.Purchase{
		ID:      uint64(p.ID.V),
		OrderID: uint64(p.ID.V),
	}
}
