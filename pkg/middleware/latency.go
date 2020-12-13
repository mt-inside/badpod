package middleware

import (
	"net/http"
	"time"

	"github.com/mt-inside/badpod/pkg/data"
)

func latencyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(data.GetDelay()) * time.Second)
		next.ServeHTTP(w, r)
	})
}
