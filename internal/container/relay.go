//go:build relay

package container

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/yael-castro/outbox/internal/app/business"
	"github.com/yael-castro/outbox/internal/app/input/command"
	mykafka "github.com/yael-castro/outbox/internal/app/output/kafka"
	"github.com/yael-castro/outbox/internal/app/output/postgres"
	"log"
	"os"
)

func New() Container {
	return new(relay)
}

type relay struct {
	container
}

func (r *relay) Inject(ctx context.Context, a any) (err error) {
	switch a := a.(type) {
	case *command.Command:
		return r.injectCommand(ctx, a)
	case **kafka.Producer:
		return r.injectProducer(ctx, a)
	}

	return r.container.Inject(ctx, a)
}

func (r *relay) injectCommand(ctx context.Context, cmd *command.Command) (err error) {
	// External dependencies
	errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	infoLogger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)

	var db *sql.DB
	if err = r.Inject(ctx, &db); err != nil {
		return
	}

	var producer *kafka.Producer
	if err = r.Inject(ctx, &producer); err != nil {
		return
	}

	// Secondary adapters
	reader := postgres.NewMessagesReader(db)
	sender := mykafka.NewMessageSender(mykafka.MessageSenderConfig{
		Info:     infoLogger,
		Error:    errLogger,
		Producer: producer,
	})
	confirmer := postgres.NewMessageDeliveryConfirmer(db)

	// Business logic
	messagesRelay := business.NewMessagesRelay(business.MessagesRelayConfig{
		Reader:    reader,
		Sender:    sender,
		Logger:    infoLogger,
		Confirmer: confirmer,
	})

	// Primary adapters
	*cmd = command.Relay(messagesRelay, errLogger)
	return
}

func (r *relay) injectProducer(_ context.Context, producer **kafka.Producer) (err error) {
	const kafkaServersEnv = "KAFKA_SERVERS"

	kafkaServers := os.Getenv(kafkaServersEnv)
	if len(kafkaServers) < 1 {
		return fmt.Errorf("missing environment variable '%s'", kafkaServersEnv)
	}

	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServers,
	})
	if err != nil {
		return
	}

	*producer = kafkaProducer
	return
}
