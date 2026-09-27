package resetpassword

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/resend/resend-go/v3"
)

const PasswordResetEmailHeader = "****Reset Password****"

func (s *service) ResetPasswordService(ctx context.Context, email string) error {
	_, err := s.repo.EmailValid(ctx, email)

	if err != nil {
		return err
	}

	token := uuid.New().String()

	if err := s.repo.SaveResetUrlInRedis(ctx, email, token, 5*time.Minute); err != nil {
		return err
	}

	errChan := make(chan error, 1)

	go func() {
		params := &resend.SendEmailRequest{
			From:    "noreply@selleasy.shop",
			To:      []string{email},
			Subject: PasswordResetEmailHeader,
			Html:    fmt.Sprintf("OTP for password reset link is  %s", s.ConstructResetLink(token)),
		}

		_, err := s.resend.Emails.Send(params)
		if err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}

}

func (s *service) ConstructResetLink(token string) string {
	return s.fEndUrl + "/reset-password?token=" + token
}
