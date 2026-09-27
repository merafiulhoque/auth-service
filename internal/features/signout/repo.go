package signout

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"context"
)

func (r *repo) DeleteRefreshTokenFromRedis(ctx context.Context, email string) error {
	redisRefreshTokenKey := domain.ConstRedisRefreshTokenKey + email
	val, err := r.rdb.Del(ctx, redisRefreshTokenKey).Result()

	if val > 0 {
		return nil
	} else if val == 0 {
		return domainerrors.ErrEntryNotFound
	}

	if err != nil {
		return err
	}
	return nil
}
