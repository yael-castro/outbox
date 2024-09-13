package business

import (
	"context"
	"errors"
	"log"
	"time"
)

type MessagesRelayConfig struct {
	Confirmer MessageDeliveryConfirmer
	Reader    MessagesReader
	Sender    MessageSender
	Logger    *log.Logger
}

func NewMessagesRelay(config MessagesRelayConfig) MessagesRelay {
	return messagesRelay{
		confirmer: config.Confirmer,
		reader:    config.Reader,
		sender:    config.Sender,
		info:      config.Logger,
	}
}

type messagesRelay struct {
	confirmer MessageDeliveryConfirmer
	reader    MessagesReader
	sender    MessageSender
	info      *log.Logger
}

func (m messagesRelay) RelayMessages(ctx context.Context) (err error) {
	var messages []Message

	for {
		messages, err = m.reader.ReadMessages(ctx)
		if err != nil {
			return
		}

		if len(messages) == 0 {
			const waitTimeForRetry = 100 * time.Millisecond

			select {
			case <-ctx.Done():
				return errors.Join(err, m.reader.Close())
			case <-time.After(waitTimeForRetry):
			}
		}

		for i := range messages {
			err = m.sender.SendMessage(ctx, &messages[i])
			if err != nil {
				continue
			}

			_ = m.confirmer.ConfirmMessageDelivery(ctx, messages[i].ID) // TODO: manage error
		}
	}
}
