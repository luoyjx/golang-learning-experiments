package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// HSET command
	err := rdb.HSet(context.Background(), "myhash", "field1", "value1").Err()
	if err != nil {
		panic(err)
	}

	// HSETNX command
	created, err := rdb.HSetNX(context.Background(), "myhash", "field2", "value2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HSETNX created new field:", created)

	// HGET command
	val, err := rdb.HGet(context.Background(), "myhash", "field1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGET field1:", val)

	// HEXISTS command
	exists, err := rdb.HExists(context.Background(), "myhash", "field2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HEXISTS field2:", exists)

	// HDEL command
	deleted, err := rdb.HDel(context.Background(), "myhash", "field1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HDEL deleted fields:", deleted)

	// HLEN command
	length, err := rdb.HLen(context.Background(), "myhash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HLEN:", length)

	// HSTRLEN command
	strLength, err := rdb.HStrLen(context.Background(), "myhash", "field2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HSTRLEN field2:", strLength)

	// HMSET command
	err = rdb.HMSet(context.Background(), "myhash", map[string]interface{}{
		"field3": "value3",
		"field4": "value4",
	}).Err()
	if err != nil {
		panic(err)
	}

	// HMGET command
	vals, err := rdb.HMGet(context.Background(), "myhash", "field2", "field3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HMGET field2 and field3:", vals)

	// HKEYS command
	keys, err := rdb.HKeys(context.Background(), "myhash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HKEYS:", keys)

	// HVALS command
	vals, err = rdb.HVals(context.Background(), "myhash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HVALS:", vals)

	// HGETALL command
	all, err := rdb.HGetAll(context.Background(), "myhash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGETALL:", all)
}
