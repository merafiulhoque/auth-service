package clients

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func CreateNewRedisClient(redisUrl string, password string, db int) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url, %w", err)
	}

	client := redis.NewClient(opt)

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect redis, %w", err)
	}
	return client, nil
}
