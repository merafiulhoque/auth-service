package middleware

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/token"
	"context"
	"net/http"
)

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := token.GetToken(r.Header.Get("Authorization"))
			if err != nil {
				response.Error(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			email, err := token.VerifyToken(tokenString, secret)

			if err != nil {
				response.Error(w, http.StatusUnauthorized, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), domain.UserEmailKey, email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
