package http

import "github.com/yael-castro/outbox/internal/app/business"

type Purchase struct {
	ID      uint64 `json:"id"`
	OrderID uint64 `json:"order_id"`
}

func (p *Purchase) ToBusiness() *business.Purchase {
	return (*business.Purchase)(p)
}
