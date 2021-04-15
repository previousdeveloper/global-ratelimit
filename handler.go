package main

import (
	"net/http"
	"net/http/httputil"
	"simple-ratelimit-service/pkg/actions"
	"simple-ratelimit-service/pkg/ratelimit"
)

type handler struct {
	proxy  *httputil.ReverseProxy
	config *ratelimit.RateLimiterConfig
	action map[string]actions.Action
}

type HttpHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewHandler(proxy *httputil.ReverseProxy,
	config *ratelimit.RateLimiterConfig,
	action map[string]actions.Action) HttpHandler {
	return &handler{proxy: proxy,
		config: config,
		action: action,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	isAllowRequest := IsAllowRequest(h.action, h.config.RateLimits, r)

	if isAllowRequest {
		h.proxy.ServeHTTP(w, r)
	} else {
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(message))
		w.Header().Set("Content-Type", "text/plain")
	}
}
