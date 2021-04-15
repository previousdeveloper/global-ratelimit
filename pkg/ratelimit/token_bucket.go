package ratelimit

import (
	"context"
	"simple-ratelimit-service/pkg/client"
	"strconv"
	"time"
)

type tokenBucketRateLimiter struct {
	redis client.RedisClient
}

func NewTokenBucket(redisClient client.RedisClient) RateLimiter {
	return &tokenBucketRateLimiter{redis: redisClient}
}

func (r *tokenBucketRateLimiter) RateLimit(ctx context.Context, genericKey string, intervalInSeconds int64, maximumRequests int64) bool { // userID can be apikey, location, ip
	value, _ := r.redis.Get(ctx, genericKey+"_last_reset_time").Result()
	lastResetTime, _ := strconv.ParseInt(value, 10, 64)
	// if the key is not available, i.e., this is the first request, lastResetTime will be set to 0 and counter be set to max requests allowed
	// check if time window since last counter reset has elapsed
	if time.Now().Unix()-lastResetTime >= intervalInSeconds {
		// if elapsed, reset the counter
		r.redis.Set(ctx, genericKey+"_counter", strconv.FormatInt(maximumRequests, 10), 0)
	} else {
		value, _ := r.redis.Get(ctx, genericKey+"_counter").Result()
		requestLeft, _ := strconv.ParseInt(value, 10, 64)
		if requestLeft <= 0 { // request left is 0 or < 0
			// drop request
			return false
		}
	}

	// decrement request count by 1
	r.redis.Decr(ctx, genericKey+"_counter")

	// handle request
	return true
}
