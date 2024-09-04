package business

import (
	"context"
	"errors"
	"log"
	"time"
)

type PurchasesNotifierConfig struct {
	Confirmer MessageDeliveryConfirmer
	Reader    PurchaseMessagesReader
	Sender    PurchaseMessageSender
	Logger    *log.Logger
}

func NewPurchasesNotifier(config PurchasesNotifierConfig) PurchasesNotifier {
	return purchasesNotifier{
		confirmer: config.Confirmer,
		reader:    config.Reader,
		sender:    config.Sender,
		info:      config.Logger,
	}
}

type purchasesNotifier struct {
	confirmer MessageDeliveryConfirmer
	reader    PurchaseMessagesReader
	sender    PurchaseMessageSender
	info      *log.Logger
}

func (p purchasesNotifier) NotifyPurchases(ctx context.Context) (err error) {
	var messages []PurchaseMessage

	for {
		messages, err = p.reader.ReadPurchaseMessages(ctx)
		if err != nil {
			return
		}

		if len(messages) == 0 {
			const waitTimeForRetry = 100 * time.Millisecond

			select {
			case <-ctx.Done():
				return errors.Join(err, p.reader.Close())
			case <-time.After(waitTimeForRetry):
			}
		}

		for _, message := range messages {
			err = p.sender.SendPurchaseMessage(ctx, message)
			if err != nil {
				continue
			}

			_ = p.confirmer.ConfirmMessageDelivery(ctx, message) // TODO: manage error
		}
	}
}
