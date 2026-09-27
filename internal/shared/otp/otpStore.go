package otp

import (
	"auth-service/internal/shared/domainerrors"
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	rdb *redis.Client
}

func NewOTPStore(rdb *redis.Client) *Store {
	return &Store{
		rdb: rdb,
	}
}

func (s *Store) VerifyAndConsumeOTP(ctx context.Context, redisKey string, userOTP string) (bool, error) {
	val, err := s.rdb.Get(ctx, redisKey).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, domainerrors.ErrOtpExpired
		}
		return false, err
	}

	if val != userOTP {
		return false, domainerrors.ErrOtpWrong
	}

	if err := s.rdb.Del(ctx, redisKey).Err(); err != nil {
		return false, err
	}

	return true, nil
}
