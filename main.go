package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Make sure to initialize the repo

func main() {
	ctx := context.Background()

	// simple get and set elements in redis.
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Print("key", val)
}
