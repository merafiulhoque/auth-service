package verifyotp

import (
	"auth-service/internal/shared/domainerrors"
	"context"
)

func (s *service) VerifyOTPService(ctx context.Context, email string, emailOtp string) error {
	ok, err := s.store.VerifyAndConsumeOTP(ctx, email, emailOtp)

	if err != nil {
		return err
	}

	if !ok {
		return domainerrors.ErrUnknownError
	}
	return nil
}
