package main

import (
	"net/http"
	"simple-ratelimit-service/pkg/actions"
	"simple-ratelimit-service/pkg/ratelimit"
)

func IsAllowRequest(actions map[string]actions.Action, policies []ratelimit.RateLimitData, r *http.Request) bool {

	var allow = false
	//TODO: Instead of loop decide one shot
	for _, rateLimit := range policies {
		allowNextAction, allowNextRequest := actions[r.Method].AllowAction(r, rateLimit)
		allow = allowNextRequest

		if !allowNextAction {
			break
		}
	}
	return allow
}
