package updatepassword

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
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
	repo *repo
}

func newService(r *repo) *service {
	return &service{
		repo: r,
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

func CreateNewHandler(db *sql.DB, rdb *redis.Client) *handler {
	return newHandler(
		newService(
			newRepo(db, rdb),
		),
	)
}
