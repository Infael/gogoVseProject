package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Infael/gogoVseProject/auth"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokens := strings.Split(authHeader, " ")

		if len(tokens) == 2 {
			cliams, err := auth.VerifyToken(tokens[1])
			if err == nil {
				// pass to next claims to next middleware
				ctx := context.WithValue(r.Context(), "email", cliams["emial"])
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		http.Error(w, "Unauthorized user.", http.StatusUnauthorized)
	})
}
