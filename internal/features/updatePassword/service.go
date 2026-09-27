package updatepassword

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) UpdatepasswordService(ctx context.Context, token string, password string) error {
	email, err := s.repo.GetEmailFromRedis(ctx, token)

	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domainerrors.ServerError
	}
	ok, err := s.repo.Updatepassword(ctx, email, string(hash))
	if err != nil {
		return err
	}

	if !ok {
		return domainerrors.ErrUnknownError
	}

	if err := s.repo.DeleteRedisEntry(ctx, domain.ConstRedisResetLinkKey+token); err != nil {
		return err
	}
	return nil
}
