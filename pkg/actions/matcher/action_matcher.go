package matcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"simple-ratelimit-service/pkg/ratelimit"
	"simple-ratelimit-service/urlpath"
	"strings"
)

type ActionMatcher interface {
}

func IsMatchedRequest(r *http.Request, limit ratelimit.RateLimitData) (string, bool) {
	var key = ""
	var isMatched = false
	action := limit.Actions
	if strings.ToLower(action.From) == "path" {
		key, isMatched = fromPath(action.Endpoint, r.URL.Path, action.Keys)
	}

	if strings.ToLower(action.From) == "query" {
		key, isMatched = fromQuery(action.Endpoint, r.URL.Path, action.Keys, r.URL)
	}

	if strings.ToLower(action.From) == "body" {
		bytes, _ := ioutil.ReadAll(r.Body)
		var body map[string]interface{}
		_ = json.Unmarshal(bytes, &body)
		key, isMatched = fromBody(action.Endpoint, r.URL.Path, action.Keys, body)
	}

	//TODO: Only header
	key += fromHeader(action.HeaderKey, r)

	return key, isMatched
}

func fromPath(endpoint, path, key string) (string, bool) {
	var genericKey = ""

	matchUrl := urlpath.New(endpoint)
	match, ok := matchUrl.Match(path)
	if !ok {
		return genericKey, false
	}

	keys := strings.Split(key, ",")
	for _, eachKey := range keys {
		genericKey += match.Params[eachKey]
	}

	return genericKey, true
}

func fromQuery(endpoint, path, key string, rawUrl *url.URL) (string, bool) {
	var genericKey = ""

	matchUrl := urlpath.New(endpoint)
	_, ok := matchUrl.Match(path)
	if !ok {
		return genericKey, false
	}

	keys := strings.Split(key, ",")
	for _, eachKey := range keys {
		genericKey += rawUrl.Query().Get(eachKey)
	}

	return genericKey, true
}

func fromBody(endpoint, path, bodyKey string, data map[string]interface{}) (string, bool) {
	var genericKey = ""

	matchUrl := urlpath.New(endpoint)
	_, ok := matchUrl.Match(path)
	if !ok {
		return genericKey, false
	}

	genericKey += data[bodyKey].(string)
	return genericKey, true
}

func fromHeader(headerKey string, r *http.Request) string {
	key := r.Header.Get(headerKey)

	if key != "" {
		return key
	}
	return ""
}
