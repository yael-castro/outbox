package command

import (
	"context"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
)

type Command func(context.Context)

func Relay(relay business.MessagesRelay, errLogger *log.Logger) Command {
	return func(ctx context.Context) {
		err := relay.RelayMessages(ctx)
		if err != nil {
			errLogger.Println(err)
		}
	}
}
