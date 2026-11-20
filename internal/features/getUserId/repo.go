package getuserid

import "context"

func (r *repo) GetUserId(ctx context.Context, email string) (int, error) {
	var id int

	err := r.db.QueryRowContext(
		ctx,
		"SELECT id FROM users WHERE email=$1",
		email,
	).Scan(&id)

	return id, err
}
