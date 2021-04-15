package actions

import (
	"net/http"
	"simple-ratelimit-service/pkg/ratelimit"
)

const (
	GetAction  = "GET"
	PostAction = "POST"
)
type Action interface {
	AllowAction(r *http.Request, limit ratelimit.RateLimitData) (bool, bool)
}

//TODO: Migrate new function
func GetActions(rateLimiter ratelimit.RateLimiter) map[string]Action {
	allAction := make(map[string]Action, 0)
	allAction[GetAction] = NewGetAction(rateLimiter)
	allAction[PostAction] = NewPostAction(rateLimiter)

	return allAction
}