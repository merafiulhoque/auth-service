package signout

import "github.com/redis/go-redis/v9"

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
	rdb *redis.Client
}

func newRepo(rdb *redis.Client) *repo {
	return &repo{
		rdb: rdb,
	}
}

func CreateNewHandler(rdb *redis.Client) *handler {
	return newHandler(
		newService(
			newRepo(rdb),
		),
	)
}
