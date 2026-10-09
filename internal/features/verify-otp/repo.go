package verifyotp

import "context"

func (r *repo) MarkEmailAsVerified(ctx context.Context, email string) error {
	_, err := r.db.ExecContext(
		ctx,
		"UPDATE users SET emailverified = $1 WHERE email = $2",
		true,
		email,
	)
	return err
}
