package signup

import (
	"auth-service/internal/shared/domainerrors"
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignupService(ctx context.Context, data *SignupRequest) (int, error) {

	exists := s.r.EmailExists(ctx, data.Email)

	if exists {
		return 0, domainerrors.ErrUserAlreadyExists
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}
	return s.r.CreateUser(ctx, data.Email, string(hash))
}
