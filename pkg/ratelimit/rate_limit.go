package ratelimit

import (
	"context"
	"simple-ratelimit-service/pkg/client"
)

const (
	TokenBucket = "token_bucket"
	FixedWindow = "fixed_window"
)

type RateLimiter interface {
	RateLimit(ctx context.Context, genericKey string, intervalInSeconds int64, maximumRequests int64) bool
}
//TODO:Init
func GetAllAlgorithms(redis client.RedisClient) map[string]RateLimiter {
	allAction := make(map[string]RateLimiter, 0)
	allAction[TokenBucket] = NewTokenBucket(redis)
	allAction[FixedWindow] = NewFixedWindow(redis)

	return allAction
}