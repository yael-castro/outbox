package kafka

import (
	"context"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
)

func NewPurchaseMessageSender(info *log.Logger) business.PurchaseMessageSender {
	return purchaseMessageSender{
		info: info,
	}
}

type purchaseMessageSender struct {
	info *log.Logger
}

func (p purchaseMessageSender) SendPurchaseMessage(ctx context.Context, purchase business.PurchaseMessage) error {
	p.info.Printf("Message: %+v\n", purchase) // TODO: implement me!
	return nil
}
