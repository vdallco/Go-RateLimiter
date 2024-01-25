# Go-RateLimiter

Simple rate limiter in Golang


# How it works

This rate limiter stores a hashmap of API keys mapped to request timestamps. When a new request is made, the rate limiter will remove any request timestamps older than n seconds. Before the request is processed, the current timestamp is added to the map for the API key. The rate limiter then enforces that the number of requests the API caller has sent in the past n seconds is equal to or less than m.


# How to run

```
go run .\rateLimiter.go
```

# Output


```
Rate limit config
10 max requests per 10 seconds

Running test for APIKeyTest123
[1706145684] [APIKeyTest123] API call # 1 successful
[1706145684] [APIKeyTest123] API call # 2 successful
[1706145684] [APIKeyTest123] API call # 3 successful
[1706145695] [APIKeyTest123] API call # 4 successful
[1706145695] [APIKeyTest123] API call # 5 successful
[1706145695] [APIKeyTest123] API call # 6 successful
[1706145695] [APIKeyTest123] API call # 7 successful
[1706145702] [APIKeyTest123] API call # 8 successful
[1706145702] [APIKeyTest123] API call # 9 successful
[1706145702] [APIKeyTest123] API call # 10 successful
[1706145702] [APIKeyTest123] API call # 11 successful
[1706145705] [APIKeyTest123] API call # 12 successful
[1706145705] [APIKeyTest123] API call # 13 successful
[1706145705] [APIKeyTest123] Error 429: Too many requests. API call # 14 rate limited

Running test for RateLimiterTest2024
[1706145705] [RateLimiterTest2024] API call # 1 successful
[1706145705] [RateLimiterTest2024] API call # 2 successful
[1706145705] [RateLimiterTest2024] API call # 3 successful
[1706145705] [RateLimiterTest2024] API call # 4 successful
[1706145705] [RateLimiterTest2024] API call # 5 successful
[1706145705] [RateLimiterTest2024] API call # 6 successful
[1706145705] [RateLimiterTest2024] API call # 7 successful
[1706145705] [RateLimiterTest2024] API call # 8 successful
[1706145705] [RateLimiterTest2024] API call # 9 successful
[1706145705] [RateLimiterTest2024] API call # 10 successful
[1706145705] [RateLimiterTest2024] Error 429: Too many requests. API call # 11 rate limited
```
