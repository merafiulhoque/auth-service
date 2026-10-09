package signin

import (
	"auth-service/internal/shared/domainerrors"
	"context"
	"database/sql"
	"errors"
	"time"
)

func (r *repo) GetUserByEmail(ctx context.Context, email string) (int, string, error) {
	var id int
	var hash string
	if err := r.db.QueryRowContext(ctx,
		"SELECT id, password FROM users WHERE email=$1", email).Scan(&id, &hash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", domainerrors.ErrInvalidCredentials
		}
		return 0, "", err
	}
	return id, hash, nil
}

func (r *repo) UpdateLastLoginTime(ctx context.Context, email string) error {
	res, err := r.db.ExecContext(ctx, "UPDATE users SET lastlogintime=$1 WHERE email=$2", time.Now(), email)

	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num <= 0 || num > 1 {
		return domainerrors.ErrUnknownError
	}
	return nil
}
