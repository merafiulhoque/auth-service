package sendotp

import (
	"context"
	"fmt"
	"time"

	"auth-service/internal/shared/domainerrors"
	"auth-service/internal/shared/utility"

	"github.com/resend/resend-go/v3"
)

const EmailVerificationHeader string = "**Email Verification**"

func (s *service) SendOTPService(ctx context.Context, data sendotp) error {
	// 1. Check if email exists in DB
	exists, err := s.r.IsEmailValid(ctx, data.Email)
	if err != nil {
		return err
	}
	if !exists {
		return domainerrors.ErrInvalidInput // Fixed: returning proper error
	}

	// 2. Generate OTP synchronously (it's fast CPU work, no goroutine needed)
	otp, err := utility.GenerateOTP()
	if err != nil {
		return err
	}

	// 3. Save to Redis first (synchronous so we fail fast if Redis is down)
	if err := s.r.SaveOTPInRedis(ctx, data.Email, otp, 1*time.Minute); err != nil {
		return err
	}

	// 4. Send email asynchronously using a buffered channel to prevent leaks
	errChan := make(chan error, 1) // Buffered channel of size 1 is safe for single-value sends
	go func() {
		params := &resend.SendEmailRequest{
			From:    "noreply@selleasy.shop",
			To:      []string{data.Email},
			Html:    fmt.Sprintf("Welcome to Our Platform. OTP for your verification is %s. Thank you.", otp),
			Subject: EmailVerificationHeader,
		}

		_, err := s.emailSender.Emails.Send(params)
		if err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	// Optionally wait for the email or let it run in background.
	// If you want to return the email error to the user:
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
