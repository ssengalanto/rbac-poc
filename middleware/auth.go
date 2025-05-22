package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("X-User-Role")
		if role == "" {
			http.Error(w, "Missing role", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "role", role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
