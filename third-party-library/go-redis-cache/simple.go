package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type User struct {
	ID   int
	Name string
}

func main() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
			"server2": ":6380",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return json.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return json.Unmarshal(b, v)
		},
	}

	user := &User{
		ID:   42,
		Name: "John Smith",
	}

	ctx := context.Background()

	if err := codec.Set(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("user:%d", user.ID),
		Value: user,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}

	var savedUser User
	if err := codec.Get(ctx, fmt.Sprintf("user:%d", user.ID), &savedUser); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", savedUser)
}
