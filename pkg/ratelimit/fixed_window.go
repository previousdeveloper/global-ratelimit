package ratelimit

import (
	"golang.org/x/net/context"
	"simple-ratelimit-service/pkg/client"
	"strconv"
	"time"
)

const expireKeyInSecond = 120

type fixedWindowRateLimiter struct {
	redis client.RedisClient
}

func NewFixedWindow(redisClient client.RedisClient) RateLimiter {
	return &fixedWindowRateLimiter{redis: redisClient}
}

func (r *fixedWindowRateLimiter) RateLimit(ctx context.Context, genericKey string, intervalInSeconds int64, maximumRequests int64) bool {
	if emptyKey(genericKey) {
		return true
	}
	i := time.Now().Unix() / intervalInSeconds
	currentWindow := strconv.FormatInt(i, 10)
	key := genericKey + ":" + currentWindow
	value, _ := r.redis.Get(ctx, key).Result()
	requestCount, _ := strconv.ParseInt(value, 10, 64)
	if overThreshold(requestCount, maximumRequests) {
		return false
	}
	var duration = time.Second * time.Duration(intervalInSeconds+expireKeyInSecond)

	go incrRequestCount(ctx, r.redis, key, duration)

	return true
}

func incrRequestCount(ctx context.Context, redis client.RedisClient, key string, duration time.Duration) {
	redis.Incr(ctx, key)
	redis.Expire(ctx, key, duration)
}

func overThreshold(requestCount int64, maximumRequests int64) bool {
	return requestCount >= maximumRequests
}

func emptyKey(genericKey string) bool {
	return genericKey == ""
}
