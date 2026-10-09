package verifyotp

import (
	"auth-service/internal/shared/domain"
	"auth-service/internal/shared/domainerrors"
	"context"
	"fmt"
)

func (s *service) VerifyOTPService(ctx context.Context, email string, emailOtp string) error {
	ok, err := s.store.VerifyAndConsumeOTP(ctx, domain.ConstRedisOtpKey+email, emailOtp)

	if err != nil {
		return err
	}

	if !ok {
		return domainerrors.ErrUnknownError
	}

	err = s.repo.MarkEmailAsVerified(ctx, email)
	if err != nil {
		return fmt.Errorf("otp verification failed, please try again after sometime...")
	}
	return nil
}
