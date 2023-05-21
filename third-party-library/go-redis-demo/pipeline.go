package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	pipe := rdb.Pipeline()

	incr := pipe.Incr(ctx, "counter")
	pipe.Expire(ctx, "counter", time.Hour)

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
