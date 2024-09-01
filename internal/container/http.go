//go:build http

package container

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
	inhttp "github.com/yael-castro/outbox/internal/app/input/http"
	"github.com/yael-castro/outbox/internal/app/output/postgres"
	"net/http"
)

func Inject(ctx context.Context, mux *http.ServeMux) (err error) {
	var db sql.DB

	if err = inject(ctx, &db); err != nil {
		return err
	}

	saver := postgres.NewPurchaseSaver(&db)
	confirmer := business.NewPurchaseConfirmer(saver)

	errFunc := inhttp.ErrorFunc(nil)
	mux.HandleFunc("POST /v1/purchases", inhttp.PostPurchase(confirmer, errFunc))
	return
}
