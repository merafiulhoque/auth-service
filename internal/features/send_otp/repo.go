package sendotp

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"context"
	"time"
)

func (r *repo) IsEmailValid(ctx context.Context, email string) (bool, error) {
	var exists bool

	if err := r.db.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)",
		email,
	).Scan(&exists); err != nil {
		return false, err
	}

	if !exists {
		return false, domainerrors.ErrUnknownError
	}

	return exists, nil
}

func (r *repo) SaveOTPInRedis(ctx context.Context, email string, otp string, time time.Duration) error {
	redisEmailOtpKey := domain.ConstRedisOtpKey + email

	err := r.rdb.Set(ctx, redisEmailOtpKey, otp, time).Err()
	return err
}
