package signin

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"auth-service/internal/shared/token"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SigninService(ctx context.Context, data *SigninRequest) (string, string, error) {
	hash, err := s.r.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return "", "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password)); err != nil {
		return "", "", domainerrors.ErrInvalidCredentials
	}
	if err := s.r.UpdateLastLoginTime(ctx, data.Email); err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := token.GenerateToken(data.Email, s.secret)
	if err != nil {
		return "", "", nil
	}

	rdbKey := domain.ConstRedisRefreshTokenKey + data.Email

	if err := s.rdb.Set(ctx, rdbKey, refreshToken, 7*24*time.Hour).Err(); err != nil {
		return "", "", nil
	}
	return accessToken, refreshToken, nil
}
