package middleware

import (
	"auth-service/internal/shared/response"
	"auth-service/internal/shared/token"
	"context"
	"net/http"
)

type UserEmail string

const UserEmailKey UserEmail = "email"

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		tokenString, err := token.GetToken(r.Header.Get("Authorization"))

		if err != nil {
			response.Error(w, http.StatusUnauthorized, err.Error())
			return
		}

		email, err := token.VerifyToken(tokenString, secret)

		if err != nil {
			response.Error(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(
			r.Context(),
			UserEmailKey,
			email,
		)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
