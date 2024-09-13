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

type Header struct {
	Key   string
	Value []byte
}

type Headers []Header

func (h *Headers) MarshalBinary() (data []byte, err error) {
	if len(*h) < 1 {
		return
	}

	headers := pb.Headers{
		Headers: make([]*pb.Header, len(*h)),
	}

	for i, header := range *h {
		headers.Headers[i] = &pb.Header{
			Key:   header.Key,
			Value: header.Value,
		}
	}

	return proto.Marshal(&headers)
}

func (h *Headers) UnmarshalBinary(data []byte) (err error) {
	var headers pb.Headers

	err = proto.Unmarshal(data, &headers)
	if err != nil {
		return
	}

	if len(headers.Headers) < 1 {
		return
	}

	*h = make([]Header, len(headers.Headers))

	for i, header := range headers.Headers {
		(*h)[i] = Header{
			Key:   header.Key,
			Value: header.Value,
		}
	}

	return
}

type Message struct {
	ID      sql.NullInt64
	Topic   sql.NullString
	Key     []byte
	Value   []byte
	Headers Headers
}

func (m *Message) ToBusiness() (message *business.Message) {
	message = &business.Message{
		ID:    uint64(m.ID.Int64),
		Topic: m.Topic.String,
		Key:   m.Key,
		Value: m.Value,
	}

	if len(m.Headers) < 1 {
		return
	}

	message.Headers = make([]business.Header, len(m.Headers))

	for i, header := range m.Headers {
		message.Headers[i] = (business.Header)(header) // TODO: create a function to replace headers
	}

	return
}
