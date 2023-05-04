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

	// create a new context
	ctx := context.Background()

	// add members to a sorted set
	if err := rdb.ZAdd(ctx, "myzset", &redis.Z{Score: 1.0, Member: "one"}, &redis.Z{Score: 2.0, Member: "two"}).Err(); err != nil {
		panic(err)
	}

	// get the number of members in a sorted set
	count, err := rdb.ZCard(ctx, "myzset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of members in myzset:", count)

	// get the score of a member in a sorted set
	score, err := rdb.ZScore(ctx, "myzset", "one").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Score of member 'one':", score)

	// get the rank of a member in a sorted set
	rank, err := rdb.ZRank(ctx, "myzset", "two").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Rank of member 'two':", rank)

	// get the members of a sorted set within a score range
	members, err := rdb.ZRangeByScore(ctx, "myzset", &redis.ZRangeBy{
		Min: "1",
		Max: "2",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members within score range [1, 2]:", members)

	// remove a member from a sorted set
	if err := rdb.ZRem(ctx, "myzset", "one").Err(); err != nil {
		panic(err)
	}
}
