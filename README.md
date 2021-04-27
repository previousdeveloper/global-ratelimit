# global-ratelimit
Distributed rate limiter, using Redis.

This package allows us to have a distributed rate limiter, using Redis or Couchbase as a central counter.

We can use it for sidecar model. Now rate limiter supports two rate limit algorithm which are token based and window rate.

## How it works


```yaml 
rate_limits:
  - actions:
       endpoint: '/product-recommendation/:id/*'
       method: GET
       keys: "id"
       requests_per_unit: 1
       unit: minute
       header_key: "test"
  - actions:
       endpoint: '/product-recommendation/*'
       method: POST
       from: "body"
       body_keys: "supplierId"
       requests_per_unit: 5
       unit: hour
```

You create config decleration, by specifying a limit and an interval. For example, maybe you want limit by using path variable or query string variable.

