package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pubsub := rdb.Subscribe(ctx, "mychannel")

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println(msg.Channel, msg.Payload)
		}
	}()

	err := rdb.Publish(ctx, "mychannel", "hello world").Err()
	if err != nil {
		log.Println(err)
	}

	time.Sleep(time.Second)
}
