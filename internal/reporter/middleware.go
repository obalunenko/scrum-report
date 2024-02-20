package reporter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	log "github.com/obalunenko/logger"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := log.FromContext(r.Context())

		ctx := log.ContextWithLogger(r.Context(), l)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
func logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		log.WithFields(ctx, log.Fields{
			"method": r.Method,
			"url":    r.URL.String(),
		}).Info("Request received")

		next.ServeHTTP(w, r)
	})
}

func logResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		ctx := r.Context()

		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r)

		log.WithFields(ctx, log.Fields{
			"method":  r.Method,
			"url":     r.URL.String(),
			"latency": time.Since(now).String(),
		}).Info("Response sent")
	})
}

type requestIDKey struct{}

func requestIDMiddleware(next http.Handler) http.Handler {
	const requestIDHeader = "X-Request-ID"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idKeyHeader := http.CanonicalHeaderKey(requestIDHeader)

		rid := r.Header.Get(idKeyHeader)

		if rid == "" {
			// New random request ID.
			rid = newRequestID()

			r.Header.Set(idKeyHeader, rid)
		}

		ctx := context.WithValue(r.Context(), requestIDKey{}, rid)

		l := log.FromContext(ctx)
		l = l.WithField("request_id", rid)

		ctx = log.ContextWithLogger(ctx, l)

		w.Header().Set(idKeyHeader, rid)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func newRequestID() string {
	u := uuid.New()

	return u.String()
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		status:         http.StatusOK,
	}
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status

	rw.ResponseWriter.WriteHeader(status)
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.WithError(r.Context(), errors.New(fmt.Sprint(err))).Error("Panic recovered")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			// Set CORS headers.
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			w.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
