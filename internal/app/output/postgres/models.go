package postgres

import (
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
	"github.com/yael-castro/outbox/pkg/pb"
	"google.golang.org/protobuf/proto"
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

func (p *Purchase) MarshalBinary() ([]byte, error) {
	purchase := pb.Purchase{
		Id:      p.ID.Int64,
		OrderId: p.OrderID.Int64,
	}

	return proto.Marshal(&purchase)
}

type Message struct {
	ID      sql.NullInt64
	Key     sql.NullString
	Topic   sql.NullString
	Header  []byte
	Content []byte
}

func (m *Message) ToBusiness() *business.Message {
	return &business.Message{
		ID:      uint64(m.ID.Int64),
		Key:     m.Key.String,
		Topic:   m.Topic.String,
		Header:  m.Header,
		Content: m.Content,
	}
}
