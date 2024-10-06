package command

import (
	"context"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
)

// Relay builds the command for relay messages
func Relay(relay business.MessagesRelay, errLogger *log.Logger) Command {
	return func(ctx context.Context) int {
		err := relay.RelayMessages(ctx)
		if err != nil {
			errLogger.Println(err)
			return fatalExitCode
		}

		return successExitCode
	}
}
