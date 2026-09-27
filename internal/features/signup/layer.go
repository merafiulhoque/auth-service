package signup

import (
	"context"
	"database/sql"
)

type handler struct {
	s *service
}
type service struct {
	r SignupRepo
}

type repo struct {
	db *sql.DB
}

type SignupRepo interface {
	EmailExists(ctx context.Context, email string) bool
	CreateUser(ctx context.Context, email string, hash string) (int, error)
}

func newHandler(s *service) *handler {
	return &handler{s: s}
}

func newService(r *repo) *service {
	return &service{r: r}
}

func newRepo(db *sql.DB) *repo {
	return &repo{db: db}
}

func CreateNewHandler(db *sql.DB) *handler {
	return newHandler(
		newService(
			newRepo(db),
		),
	)
}
