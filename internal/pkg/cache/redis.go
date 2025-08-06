package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

//todo сделать generic

// RedisClient .
type RedisClient interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (any, error)
}

type redisClient struct {
	cl *redis.Client
}

func NewRedisClient(addr string, password string, db int) RedisClient {
	return &redisClient{
		cl: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
}

func (c *redisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return c.cl.Set(ctx, key, value, expiration).Err()
}

func (c *redisClient) Get(ctx context.Context, key string) (any, error) {
	return c.cl.Get(ctx, key).Result()
}
