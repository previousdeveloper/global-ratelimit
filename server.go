package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"simple-ratelimit-service/config"
	"simple-ratelimit-service/pkg/actions"
	"simple-ratelimit-service/pkg/client"
	_ "simple-ratelimit-service/pkg/helper"
	"simple-ratelimit-service/pkg/ratelimit"
)

const (
	message    = "too many request!"
	statusCode = http.StatusTooManyRequests
)

func main() {
	//proxyUrl := os.Getenv("http://stageproductrecommendationapi.trendyol.com")
	url, err := url.Parse("http://stageproductrecommendationapi.trendyol.com")
	if err != nil {
		panic(err)
	}

	conf, _ := ratelimit.ReadConf("rate-limit.yml")

	director := func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.Host = url.Host
	}

	redisConfig := &config.RedisConfig{URL: "localhost:6379"}
	driver := client.RedisDriver(redisConfig)
	algorithms := ratelimit.GetAllAlgorithms(driver)
	reverseProxy := &httputil.ReverseProxy{Director: director}
	handler := NewHandler(reverseProxy, conf, actions.GetActions(algorithms[ratelimit.FixedWindow]))
	http.Handle("/", handler)

	err = http.ListenAndServe(":3083", nil)
	//TODO: Json logging
	fmt.Println(err)
}
