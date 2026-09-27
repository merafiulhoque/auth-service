package clients

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func CreateNewRedisClient(redisUrl string, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("error connecting redis %v\n", err)
	}

	return rdb, nil
}
