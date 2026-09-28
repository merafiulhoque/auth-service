package middleware

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/response"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func RateLimiter(rdb *redis.Client, limit int, window time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			clientIP := r.Header.Get("X-Forwarded-For")
			if clientIP == "" {
				var err error
				clientIP, _, err = net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					clientIP = r.RemoteAddr
				}
			}
			redisKey := domain.ConstRateLimitKey + clientIP + r.URL.Path

			pipe := rdb.Pipeline()
			incr := pipe.Incr(ctx, redisKey)
			pipe.Expire(ctx, redisKey, window)

			_, err := pipe.Exec(ctx)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			if incr.Val() > int64(limit) {
				w.Header().Set("Retry-After", fmt.Sprintf("%f seconds", window.Seconds()))
				response.Error(w, http.StatusTooManyRequests, "too many requests")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
