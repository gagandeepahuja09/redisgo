package main

import (
	"context"
	"fmt"
	"time"

	rateLimiter "github.com/gagandeepahuja09/redisgo/pkg/rate_limiter"
	"github.com/gagandeepahuja09/redisgo/pkg/redisClient"
)

func main() {
	ctx := context.Background()

	redisClient.Setup()

	err := redisClient.Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	for i := 0; i < 5; i++ {
		if rateLimiter.LimitReachedRedis() {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
	time.Sleep(time.Second * 1)
	for i := 0; i < 3; i++ {
		if rateLimiter.LimitReachedRedis() {
			fmt.Println("YES1")
		} else {
			fmt.Println("NO1")
		}
	}
	time.Sleep(time.Second * 10)
	for i := 0; i <= 10; i++ {
		if rateLimiter.LimitReachedRedis() {
			fmt.Println("YES2")
		} else {
			fmt.Println("NO2")
		}
	}
}
