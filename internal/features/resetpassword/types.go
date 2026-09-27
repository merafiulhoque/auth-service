package resetpassword

type ResetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}
