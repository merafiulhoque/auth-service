package updatepassword

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"context"
	"database/sql"
	"errors"

	"github.com/redis/go-redis/v9"
)

func (r *repo) GetEmailFromRedis(ctx context.Context, token string) (string, error) {
	redisKey := domain.ConstRedisResetLinkKey + token
	email, err := r.rdb.Get(ctx, redisKey).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", domainerrors.RedisErrResetLinkExpired
		}
		return "", err
	}

	if email == "" {
		if err := r.DeleteRedisEntry(ctx, redisKey); err != nil {
			return "", err
		}
		return "", domainerrors.RedisErrResetLinkExpired
	}

	return email, nil
}

func (r *repo) Updatepassword(ctx context.Context, email string, hash string) (bool, error) {
	result, err := r.db.ExecContext(
		ctx,
		"UPDATE users SET password=$1 WHERE email=$2",
		hash,
		email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, domainerrors.ErrUserNotFound
		}
		return false, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if count != 1 {
		return false, domainerrors.ErrUnknownError
	}
	return true, nil
}

func (r *repo) DeleteRedisEntry(ctx context.Context, key string) error {
	count, err := r.rdb.Del(ctx, key).Result()

	if err != nil {
		return err
	}

	if count == 0 {
		return domainerrors.ErrUnknownError
	}
	return nil
}
