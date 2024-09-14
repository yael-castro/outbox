package http

import (
	"encoding/json"
	"errors"
	"github.com/yael-castro/outbox/internal/app/business"
	"net/http"
)

func DefaultErrorFunc(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	//goland:noinspection ALL
	switch err.(type) {
	case *json.InvalidUnmarshalError, *json.SyntaxError:
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`{"error": "incorrect json syntax"}`))
	}

	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(`{"error": "internal server error"}`))
}

func ErrorFunc(errFunc func(w http.ResponseWriter, err error)) func(http.ResponseWriter, error) {
	if errFunc == nil {
		errFunc = DefaultErrorFunc
	}

	return func(w http.ResponseWriter, err error) {
		var businessErr business.Error

		if !errors.As(err, &businessErr) {
			errFunc(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		//goland:noinspection ALL
		switch businessErr {
		case // Bad requests
			business.ErrMissingPurchaseID,
			business.ErrMissingPurchaseOrderID:
			w.WriteHeader(http.StatusBadRequest)
		case
			business.ErrDuplicatedOrderID:
			w.WriteHeader(http.StatusConflict)
		default: // Unhandled errors
			err = errors.New("unhandled error")
			w.WriteHeader(http.StatusInternalServerError)
		}

		_ = json.NewEncoder(w).Encode(map[string]string{
			"code":  businessErr.Error(),
			"error": err.Error(),
		})
	}
}
