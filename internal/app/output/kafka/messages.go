package kafka

import (
	"context"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
)

func NewMessageSender(info *log.Logger) business.MessageSender {
	return messageSender{
		info: info,
	}
}

type messageSender struct {
	info *log.Logger
}

func (p messageSender) SendMessage(_ context.Context, message *business.Message) error {
	p.info.Printf("Message: %+v\n", message) // TODO: implement me!
	return nil
}
