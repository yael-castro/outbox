package http

import (
	"encoding/json"
	"github.com/yael-castro/outbox/internal/app/business"
	"net/http"
)

func PostPurchase(c business.PurchaseConfirmer, errFunc func(http.ResponseWriter, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var purchase Purchase

		err := json.NewDecoder(r.Body).Decode(&purchase)
		if err != nil {
			errFunc(w, err)
			return
		}

		err = c.ConfirmPurchase(ctx, purchase.ToBusiness())
		if err != nil {
			errFunc(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(purchase)
	}
}
