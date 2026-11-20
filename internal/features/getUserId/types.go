package getuserid

type GetUserId struct {
	Email string `json:"email" validate:"required,email"`
}
