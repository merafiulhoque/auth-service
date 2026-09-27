package refresh

import (
	"auth-service/internal/shared/domainerrors"
	"auth-service/internal/shared/token"
	"context"
	"log/slog"
	"time"
)

func (s *service) RefreshService(ctx context.Context, refreshToken string) (string, string, error) {
	// get email from refresh token
	email, err := token.VerifyToken(refreshToken, s.secret)

	if err != nil {
		return "", "", err
	}

	//redis lookup for refresh token
	valid, err := s.repo.IsRefreshTokenValid(ctx, email, refreshToken)

	if err != nil {
		return "", "", err
	}

	if !valid {
		return "", "", domainerrors.TokenErrInvalid
	}

	accessToken, newRefreshToken, err := token.GenerateToken(email, s.secret)

	slog.Info("Tokens Issued: ", "accessToken", accessToken, "newRefreshToken", newRefreshToken)

	if err != nil {
		return "", "", err
	}

	if accessToken == "" || newRefreshToken == "" {
		return "", "", domainerrors.ServerError
	}

	if err := s.repo.UpdateRefreshTokenInRedis(ctx, email, newRefreshToken, 10*time.Minute); err != nil {
		return "", "", err
	}
	return accessToken, newRefreshToken, nil
}
