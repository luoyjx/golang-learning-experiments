package main

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{})

	p := map[string]interface{}{
		"event":    "Comment",
		"data":     "asdasd",
		"pushGuid": "asdasdasd",
		"users":    []int64{1, 2, 3},
	}

	bs, _ := json.Marshal(p)

	redisClient.Publish(context.Background(), "MODEL", bs)
}
