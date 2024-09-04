//go:build http

package container

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
	inhttp "github.com/yael-castro/outbox/internal/app/input/http"
	"github.com/yael-castro/outbox/internal/app/output/postgres"
	"log"
	"net/http"
	"os"
)

func Inject(ctx context.Context, mux *http.ServeMux) (err error) {
	var db sql.DB

	if err = inject(ctx, &db); err != nil {
		return err
	}

	errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	saver := postgres.NewPurchaseSaver(&db, errLogger)
	confirmer := business.NewPurchaseConfirmer(saver)

	errFunc := inhttp.ErrorFunc(nil)

	// Setting routes
	// TODO: should this be here?
	mux.HandleFunc("POST /v1/purchases", inhttp.PostPurchase(confirmer, errFunc))
	return
}
