package signin

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type handler struct {
	s *service
}

type service struct {
	r      SigninRepo
	secret string
	rdb    *redis.Client
}

type repo struct {
	db *sql.DB
}

type SigninRepo interface {
	GetUserByEmail(ctx context.Context, email string) (string, error)
	UpdateLastLoginTime(ctx context.Context, email string) error
}

func newHandler(s *service) *handler {
	return &handler{s: s}
}

func newService(r *repo, secret string, rdb *redis.Client) *service {
	return &service{
		r:      r,
		secret: secret,
		rdb:    rdb,
	}
}

func newRepo(db *sql.DB) *repo {
	return &repo{db: db}
}

func CreateNewHandler(db *sql.DB, secret string, rdb *redis.Client) *handler {
	return newHandler(
		newService(
			newRepo(db),
			secret,
			rdb,
		),
	)
}
