package refresh

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
	repo   *repo
	secret string
}

func newService(r *repo, s string) *service {
	return &service{
		repo:   r,
		secret: s,
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

func CreateNewHandler(rdb *redis.Client, secret string) *handler {
	return newHandler(
		newService(
			newRepo(rdb),
			secret,
		),
	)
}
