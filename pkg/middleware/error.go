package middleware

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/mt-inside/badpod/pkg/data"
)

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if data.GetErrorRate() < rand.Float64() {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(500)
			fmt.Fprintf(w, "error")
		}
	})
}
