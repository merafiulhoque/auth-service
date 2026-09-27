package signup

import (
	"context"
)

func (r *repo) CreateUser(ctx context.Context, email string, hash string) (int, error) {
	var id int

	if err := r.db.QueryRowContext(
		ctx,
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		email, hash,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) EmailExists(ctx context.Context, email string) bool {
	var exists bool

	_ = r.db.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email,
	).Scan(&exists)

	return exists
}
