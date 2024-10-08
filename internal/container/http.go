//go:build http

package container

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
	inhttp "github.com/yael-castro/outbox/internal/app/input/http"
	"github.com/yael-castro/outbox/internal/app/output/postgres"
	"github.com/yael-castro/outbox/pkg/middleware"
	"log"
	"net/http"
	"os"
)

func New() Container {
	return new(handler)
}

type handler struct {
	container
}

func (h *handler) Inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case *http.Handler:
		return h.injectHandler(ctx, a)
	}

	return h.container.Inject(ctx, a)
}

func (h *handler) injectHandler(ctx context.Context, handler *http.Handler) (err error) {
	// External dependencies
	var db *sql.DB

	if err = h.Inject(ctx, &db); err != nil {
		return err
	}

	infoLogger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	// Secondary adapters
	const purchaseTopic = "purchases"
	saver := postgres.NewPurchaseSaver(postgres.PurchaseSaverConfig{
		DB:     db,
		Topic:  purchaseTopic,
		Logger: errLogger,
	})

	// Business logic
	confirmer := business.NewPurchaseConfirmer(saver)

	// Building mux
	mux := http.NewServeMux()
	errFunc := inhttp.ErrorFunc(nil)

	// Setting routes
	// TODO: should this be here?
	mux.HandleFunc("POST /v1/purchases", inhttp.PostPurchase(confirmer, errFunc))

	// Initializing http.Handler
	*handler = middleware.Logger(mux, infoLogger)
	return
}
