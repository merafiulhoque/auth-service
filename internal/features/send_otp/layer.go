package sendotp

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/resend/resend-go/v3"
)

type handler struct {
	s *service
}
type service struct {
	r           *repo
	emailSender *resend.Client
}

type repo struct {
	db  *sql.DB
	rdb *redis.Client
}

type SignupRepo interface {
	EmailExists(ctx context.Context, email string) bool
	CreateUser(ctx context.Context, email string, hash string) (int, error)
}

func newHandler(s *service) *handler {
	return &handler{s: s}
}

func newService(r *repo, emailSender *resend.Client) *service {
	return &service{
		r:           r,
		emailSender: emailSender,
	}
}

func newRepo(db *sql.DB, rdb *redis.Client) *repo {
	return &repo{
		db:  db,
		rdb: rdb,
	}
}

func CreateNewHandler(db *sql.DB, emailSender *resend.Client, rdb *redis.Client) *handler {
	return newHandler(
		newService(
			newRepo(db, rdb),
			emailSender,
		),
	)
}
