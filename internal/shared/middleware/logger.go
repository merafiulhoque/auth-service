package middleware

import (
	"log/slog"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		method := r.Method
		ip := r.Header.Get("X-Forwarded-For")

		if ip == "" {
			ip = r.RemoteAddr
		}
		endpoint := r.URL.Path

		slog.Info("HTTP : ", "ip", ip, "method", method, "endpoint", endpoint)
		next.ServeHTTP(w, r)
	})
}
