package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Infael/gogoVseProject/service/auth"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokens := strings.Split(authHeader, " ")

		if len(tokens) == 2 {
			claims, err := auth.VerifyToken(tokens[1])
			if err == nil {
				// pass claims to the next middleware
				ctx := context.WithValue(r.Context(), "email", claims["email"])
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		http.Error(w, "Unauthorized user.", http.StatusUnauthorized)
	})
}
