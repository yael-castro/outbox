//go:build relay

package container

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
	"github.com/yael-castro/outbox/internal/app/input/command"
	"github.com/yael-castro/outbox/internal/app/output/kafka"
	"github.com/yael-castro/outbox/internal/app/output/postgres"
	"log"
	"os"
)

func Inject(ctx context.Context, cmd *command.Command) (err error) {
	// External dependencies
	errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	infoLogger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)

	var db sql.DB

	err = inject(ctx, &db)
	if err != nil {
		return
	}

	// Secondary adapters
	reader := postgres.NewPurchaseMessagesReader(&db)
	sender := kafka.NewPurchaseMessageSender(infoLogger)
	confirmer := postgres.NewMessageDeliveryConfirmer(&db)

	// Business logic
	notifier := business.NewPurchasesNotifier(business.PurchasesNotifierConfig{
		Reader:    reader,
		Sender:    sender,
		Logger:    infoLogger,
		Confirmer: confirmer,
	})

	// Primary adapters
	*cmd = command.NotifyPurchases(notifier, infoLogger, errLogger)
	return
}
