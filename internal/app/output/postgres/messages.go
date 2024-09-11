//go:build relay || tests

package postgres

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
)

func NewMessagesReader(db *sql.DB) business.MessagesReader {
	return messageReader{
		db: db,
	}
}

type messageReader struct {
	db *sql.DB
}

func (p messageReader) ReadMessages(ctx context.Context) ([]business.Message, error) {
	const defaultLimit = 100

	rows, err := p.db.QueryContext(ctx, selectPurchaseMessages, defaultLimit)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	const expectedMessages = 100
	messages := make([]business.Message, 0, expectedMessages)

	for rows.Next() {
		var rawHeaders []byte
		message := Message{}

		err = rows.Scan(
			&message.ID,
			&message.Topic,
			&message.Key,
			&rawHeaders,
			&message.Value,
		)
		if err != nil {
			return nil, err
		}

		err = message.Headers.UnmarshalBinary(rawHeaders)
		if err != nil {
			return nil, err
		}

		messages = append(messages, *message.ToBusiness())
	}

	return messages, nil
}

func (messageReader) Close() error {
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

func (m messageDeliveryConfirmer) ConfirmMessageDelivery(ctx context.Context, messageID uint64) error {
	_, err := m.db.ExecContext(ctx, updatePurchaseMessage, messageID)
	return err
}
