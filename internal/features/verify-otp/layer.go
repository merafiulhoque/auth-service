package verifyotp

import (
	"auth-service/internal/shared/otp"
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
	store *otp.Store
}

func newService(store *otp.Store) *service {
	return &service{
		store: store,
	}
}

func CreateNewHandler(db *sql.DB, rdb *redis.Client, store *otp.Store) *handler {
	return newHandler(
		newService(
			store,
		),
	)
}
