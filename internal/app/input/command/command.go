package command

import (
	"context"
	"github.com/yael-castro/outbox/internal/app/business"
	"log"
)

type Command func(context.Context)

func NotifyPurchases(notifier business.PurchasesNotifier, infoLogger, errLogger *log.Logger) Command {
	return func(ctx context.Context) {
		err := notifier.NotifyPurchases(ctx)
		if err != nil {
			errLogger.Println(err)
		}
	}
}
