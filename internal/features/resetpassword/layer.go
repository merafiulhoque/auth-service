package resetpassword

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/resend/resend-go/v3"
)

type handler struct {
	service *service
}

func newHandler(s *service) *handler {
	return &handler{
		service: s,
	}
}

type service struct {
	repo    *repo
	resend  *resend.Client
	fEndUrl string
}

func newService(r *repo, resend *resend.Client, fEndUrl string) *service {
	return &service{
		repo:    r,
		resend:  resend,
		fEndUrl: fEndUrl,
	}
}

type repo struct {
	db  *sql.DB
	rdb *redis.Client
}

func newRepo(db *sql.DB, rdb *redis.Client) *repo {
	return &repo{
		db:  db,
		rdb: rdb,
	}
}

func CreateNewHandler(db *sql.DB, rdb *redis.Client, resend *resend.Client, fEndUrl string) *handler {
	return newHandler(
		newService(
			newRepo(db, rdb),
			resend,
			fEndUrl,
		),
	)
}
