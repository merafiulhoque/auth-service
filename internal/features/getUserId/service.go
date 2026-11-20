package getuserid

import "context"

func (s *service) getUserIdService(ctx context.Context, email string) (int, error) {
	return s.repo.GetUserId(ctx, email)
}
