package client

import (
	"context"
	"github.com/go-redis/redis/v8"
	"simple-ratelimit-service/config"
	"time"
)

type redisClient struct {
	redis *redis.Client
}

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
	Decr(ctx context.Context, key string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

func (r *redisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.redis.Get(ctx, key)
}

func (r *redisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	return r.redis.Incr(ctx, key)
}

func (r *redisClient) Decr(ctx context.Context, key string) *redis.IntCmd {
	return r.redis.Decr(ctx, key)
}

func (r *redisClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return r.redis.Expire(ctx, key, expiration)
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.redis.Set(ctx, key, value, expiration)
}

func RedisDriver(config *config.RedisConfig) RedisClient {
	return &redisClient{redis: redisDriver(config)}
}

//TODO:Configuration logging (Pooling)
func redisDriver(config *config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        config.URL,
		Password:    "",
		DB:          0,
		ReadTimeout: time.Millisecond * 50,
	})
	return client
}
