package http

import (
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

		switch businessErr {
		case // Bad requests
			business.ErrMissingPurchaseID,
			business.ErrMissingPurchaseOrderID:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
