package signout

import (
	"context"
)

func (s *service) SignoutService(ctx context.Context, email string) error {
	return s.repo.DeleteRefreshTokenFromRedis(ctx, email)
}
