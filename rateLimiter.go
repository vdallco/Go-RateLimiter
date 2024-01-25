package main

import (
  "fmt"
  "time"
)

var requests = make(map[string][]int64)
var maxRequestsPerUnit = 10
var unitSeconds = 10

func rateLimiter(apiKey string) bool {
	// lookup requests
	requestsForKey := requests[apiKey]

	nowSeconds := time.Now().Unix()

	// remove old requests
	for x:=0; x<len(requestsForKey); x++ {
		if requestsForKey[x] < nowSeconds - int64(unitSeconds) {
			requestsForKey = append(requestsForKey[:x], requestsForKey[x+1:]...)
		}
	}

	// add new request
	requestsForKey = append(requestsForKey, nowSeconds)

	// update hashmap
	requests[apiKey] = requestsForKey

	// check if rate limit is exceeded
	if len(requests[apiKey]) > maxRequestsPerUnit {
		return false
	}

	return true
}

func main() {
	apiKey1 := "APIKeyTest123"
	apiKey2 := "RateLimiterTest2024"

	fmt.Println("Rate limit config")
	fmt.Printf("%d max requests per %d seconds\n\n", maxRequestsPerUnit, unitSeconds)

	fmt.Printf("Running test for %s\n", apiKey1)
	for i:=1; i<100; i++ {
		if i%4 == 0 && i<14 {
			time.Sleep(time.Duration(15-i) * time.Second)
		}
		apiCallSuccess := rateLimiter(apiKey1)
		if apiCallSuccess == false{
			fmt.Printf("[%d] [%s] Error 429: Too many requests. API call # %d rate limited", time.Now().Unix(), apiKey1, i)
			break
		} else {
			fmt.Printf("[%d] [%s] API call # %d successful\n", time.Now().Unix(), apiKey1, i)
		}
	}

	fmt.Printf("\n\nRunning test for %s\n", apiKey2)
	for i:=1; i<100; i++ {
		apiCallSuccess := rateLimiter(apiKey2)
		if apiCallSuccess == false{
			fmt.Printf("[%d] [%s] Error 429: Too many requests. API call # %d rate limited", time.Now().Unix(), apiKey2, i)
			break
		} else {
			fmt.Printf("[%d] [%s] API call # %d successful\n", time.Now().Unix(), apiKey2, i)
		}
	}

}