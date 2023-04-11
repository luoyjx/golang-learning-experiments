package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// set a key-value pair
	err := rdb.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// set a key-value pair with expiration time
	err = rdb.SetEX(context.Background(), "key", "value", time.Hour).Err()
	if err != nil {
		panic(err)
	}

	// set a key-value pair only if the key does not exist
	err = rdb.SetNX(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// set a key-value pair with expiration time in milliseconds
	err = rdb.PSetEX(context.Background(), "key", "value", time.Millisecond*500).Err()
	if err != nil {
		panic(err)
	}

	// get the value of a key
	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// get the value of a key and set a new value
	val, err = rdb.GetSet(context.Background(), "key", "new_value").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// get the length of a string value
	length, err := rdb.StrLen(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("length", length)

	// append a string value to an existing key
	err = rdb.Append(context.Background(), "key", "_suffix").Err()
	if err != nil {
		panic(err)
	}

	// increment the value of a key by 1
	err = rdb.Incr(context.Background(), "counter").Err()
	if err != nil {
		panic(err)
	}

	// increment the value of a key by a specified amount
	err = rdb.IncrBy(context.Background(), "counter", 10).Err()
	if err != nil {
		panic(err)
	}

	// decrement the value of a key by 1
	err = rdb.Decr(context.Background(), "counter").Err()
	if err != nil {
		panic(err)
	}

	// decrement the value of a key by a specified amount
	err = rdb.DecrBy(context.Background(), "counter", 5).Err()
	if err != nil {
		panic(err)
	}

	// set multiple key-value pairs
	err = rdb.MSet(context.Background(), "key1", "value1", "key2", "value2").Err()
	if err != nil {
		panic(err)
	}

	// get the values of multiple keys
	vals, err := rdb.MGet(context.Background(), "key1", "key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", vals[0], "key2", vals[1])

	// set multiple key-value pairs only if none of the keys exist
	err = rdb.MSetNX(context.Background(), "key3", "value3", "key4", "value4").Err()
	if err != nil {
		panic(err)
	}
}
