package actions

import (
	"net/http"
	"simple-ratelimit-service/pkg/actions/matcher"
	"simple-ratelimit-service/pkg/helper"
	"simple-ratelimit-service/pkg/ratelimit"
)

type getAction struct {
	rateLimiter ratelimit.RateLimiter
}

func NewGetAction(rateLimiter ratelimit.RateLimiter) Action {
	return &getAction{rateLimiter: rateLimiter}
}

func (g *getAction) AllowAction(r *http.Request, limit ratelimit.RateLimitData) (bool, bool) {

	key, isMatched := matcher.IsMatchedRequest(r, limit)

	if isMatched {
		allowNextRequest := g.rateLimiter.RateLimit(r.Context(), key, helper.GetValueByKey(limit.Actions.Unit), limit.Actions.RequestsPerUnit)
		return !isMatched, allowNextRequest
	}
	return true, true
}
