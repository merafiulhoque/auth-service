package getuserid

import "database/sql"

type handler struct {
	service *service
}

type service struct {
	repo *repo
}

type repo struct {
	db *sql.DB
}

func newHandler(s *service) *handler {
	return &handler{service: s}
}

func newService(r *repo) *service {
	return &service{repo: r}
}

func newRepo(db *sql.DB) *repo {
	return &repo{db: db}
}

func CreateNewHandler(db *sql.DB) *handler {
	return newHandler(newService(newRepo(db)))
}
