package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
)

type MessageSenderConfig struct {
	Producer *kafka.Producer
	Info     *log.Logger
	Error    *log.Logger
}

func NewMessageSender(config MessageSenderConfig) business.MessageSender {
	return messageSender{
		producer: config.Producer,
		error:    config.Error,
		info:     config.Info,
	}
}

type messageSender struct {
	producer *kafka.Producer
	info     *log.Logger
	error    *log.Logger
}

func (p messageSender) SendMessage(ctx context.Context, msg *business.Message) (err error) {
	message, err := NewMessage(msg)
	if err != nil {
		return
	}

	err = p.producer.Produce(message, nil)
	if err != nil {
		p.error.Printf("Failed to send message to Kafka topic %s: %v", msg.Topic, err)
		return
	}

	const waitTime = 1_000
	for p.producer.Flush(waitTime) > 1 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}

	p.info.Printf("Message: %+v\n", msg)
	return nil
}
