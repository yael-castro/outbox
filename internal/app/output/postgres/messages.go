//go:build relay || tests

package postgres

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
)

func NewPurchaseMessagesReader(db *sql.DB) business.PurchaseMessagesReader {
	return purchaseMessageReader{
		db: db,
	}
}

type purchaseMessageReader struct {
	db *sql.DB
}

func (p purchaseMessageReader) ReadPurchaseMessages(ctx context.Context) ([]business.PurchaseMessage, error) {
	const defaultLimit = 100

	rows, err := p.db.QueryContext(ctx, selectPurchaseMessages, defaultLimit)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	const expectedMessages = 100
	messages := make([]business.PurchaseMessage, 0, expectedMessages)

	for rows.Next() {
		message := PurchaseMessage{}

		err = rows.Scan(
			&message.ID,
			&message.Purchase.ID,
			&message.Purchase.OrderID,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message.ToBusiness())
	}

	return messages, nil
}

func (purchaseMessageReader) Close() error {
	return nil
}

func NewMessageDeliveryConfirmer(db *sql.DB) business.MessageDeliveryConfirmer {
	return messageDeliveryConfirmer{
		db: db,
	}
}

type messageDeliveryConfirmer struct {
	db *sql.DB
}

func (m messageDeliveryConfirmer) ConfirmMessageDelivery(ctx context.Context, message business.PurchaseMessage) error {
	_, err := m.db.ExecContext(ctx, updatePurchaseMessage, message.ID)
	return err
}
