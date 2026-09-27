package verifyotp

type VerifyOTP struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp" validate:"required,len=6,numeric"`
}
