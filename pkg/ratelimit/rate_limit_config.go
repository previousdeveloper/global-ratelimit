package ratelimit

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RateLimiterConfig struct {
	RateLimits []RateLimitData `yaml:"rate_limits"`
}

type RateLimitData struct {
	Actions Actions `yaml:"actions"`
}

type Actions struct {
	Endpoint        string `yaml:"endpoint"`
	Method          string `yaml:"method"`
	RequestsPerUnit int64  `yaml:"requests_per_unit"`
	Unit            string `yaml:"unit"`
	HeaderKey       string `yaml:"header_key"`
	From            string `yaml:"from"`
	Keys            string `yaml:"keys"`
	BodyKeys 		string `yaml:"body_keys"`
}

func ReadConf(filename string) (*RateLimiterConfig, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &RateLimiterConfig{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}
