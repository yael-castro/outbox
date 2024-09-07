package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(handler http.Handler, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cw := customResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		now := time.Now()

		handler.ServeHTTP(&cw, r)

		logger.Printf("%s %s %d %s\n", r.Method, r.URL, cw.statusCode, time.Now().Sub(now))
	}
}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *customResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
