package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	publish()
}

func newClient() *redis.Client {
	// use default redis options
	return redis.NewClient(&redis.Options{})
}

func getSet() {
	redisClient := newClient()

	cmd := redisClient.Set(context.Background(), "key", "value", 0)
	if cmd.Err() != nil {
		panic(fmt.Errorf("set key error: %w", cmd.Err()))
	}

	getCmd := redisClient.Get(context.Background(), "key")
	if getCmd.Err() != nil {
		if getCmd.Err() == redis.Nil {
			fmt.Println("key not found")
			return
		}

		panic(fmt.Errorf("get key error: %w", getCmd.Err()))
	}

	fmt.Println("get value by key: ", getCmd.Val())
}

func publish() {
	redisClient := newClient()

	p := map[string]interface{}{
		"event":    "Comment",
		"data":     "asdasd",
		"pushGuid": "asdasdasd",
		"users":    []int64{1, 2, 3},
	}

	bs, _ := json.Marshal(p)

	redisClient.Publish(context.Background(), "MODEL", bs)
}
