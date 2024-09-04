package http

import (
	"encoding/json"
	"errors"
	"github.com/yael-castro/outbox/internal/app/business"
	"net/http"
)

func ErrorFunc(errFunc func(w http.ResponseWriter, err error)) func(http.ResponseWriter, error) {
	if errFunc == nil {
		errFunc = func(w http.ResponseWriter, err error) {
			http.Error(w, "Internal error!", http.StatusInternalServerError)
		}
	}

	return func(w http.ResponseWriter, err error) {
		var businessErr business.Error

		if !errors.As(err, &businessErr) {
			errFunc(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		switch businessErr {
		case // Bad requests
			business.ErrMissingPurchaseID,
			business.ErrMissingPurchaseOrderID:

			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"code":  businessErr.Error(),
				"error": err.Error(),
			})
		}
	}
}
