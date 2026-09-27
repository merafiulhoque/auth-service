package resetpassword

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"context"
	"database/sql"
	"errors"
	"time"
)

func (r *repo) EmailValid(ctx context.Context, email string) (bool, error) {
	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",
		email,
	).Scan(&exists)

	if errors.Is(err, sql.ErrNoRows) {
		return false, domainerrors.ErrUserNotFound
	}

	if err != nil {
		return false, err
	}

	if !exists {
		return false, domainerrors.ErrUnknownError
	}

	return true, nil
}

func (r *repo) SaveResetUrlInRedis(ctx context.Context, email string, token string, ttl time.Duration) error {
	redisKey := domain.ConstRedisResetLinkKey + token
	_, err := r.rdb.Set(
		ctx,
		redisKey,
		email,
		ttl,
	).Result()

	return err
}
