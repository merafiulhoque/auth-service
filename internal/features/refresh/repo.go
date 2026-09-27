package refresh

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"context"
	"time"
)

func (r *repo) IsRefreshTokenValid(ctx context.Context, email string, refreshToken string) (bool, error) {
	redisKey := domain.ConstRedisRefreshTokenKey + email

	val, err := r.rdb.Get(ctx, redisKey).Result()

	if err != nil {
		return false, err
	}

	if val != refreshToken {
		return false, domainerrors.TokenErrInvalid
	}

	return true, nil
}

func (r *repo) UpdateRefreshTokenInRedis(ctx context.Context, email string, refreshToken string, ttl time.Duration) error {
	refreshTokenKey := domain.ConstRedisRefreshTokenKey + email
	_, err := r.rdb.Set(ctx, refreshTokenKey, refreshToken, ttl).Result()
	return err
}
