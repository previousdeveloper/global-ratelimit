package actions

import (
	"net/http"
	"simple-ratelimit-service/pkg/actions/matcher"
	"simple-ratelimit-service/pkg/helper"
	"simple-ratelimit-service/pkg/ratelimit"
)

type postAction struct {
	rateLimiter ratelimit.RateLimiter
}

func NewPostAction(rateLimiter ratelimit.RateLimiter) Action {
	return &postAction{rateLimiter: rateLimiter}
}

func (p *postAction) AllowAction(r *http.Request, limit ratelimit.RateLimitData) (bool, bool) {
	key, isMatched := matcher.IsMatchedRequest(r, limit)
	if isMatched {
		allowNextRequest := p.rateLimiter.RateLimit(r.Context(), key, helper.GetValueByKey(limit.Actions.Unit), limit.Actions.RequestsPerUnit)
		return !isMatched, allowNextRequest
	}
	return true, true
}
