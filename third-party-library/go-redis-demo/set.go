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

	// create a new context
	ctx := context.Background()

	// add a member to a set
	err := rdb.SAdd(ctx, "myset", "member1").Err()
	if err != nil {
		panic(err)
	}

	// add multiple members to a set
	err = rdb.SAdd(ctx, "myset", "member2", "member3").Err()
	if err != nil {
		panic(err)
	}

	// get all members of a set
	members, err := rdb.SMembers(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members of myset:")
	for _, member := range members {
		fmt.Println(member)
	}

	// check if a member exists in a set
	exists, err := rdb.SIsMember(ctx, "myset", "member1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Does member1 exist in myset?", exists)

	// remove a member from a set
	err = rdb.SRem(ctx, "myset", "member1").Err()
	if err != nil {
		panic(err)
	}

	// get all members of a set again
	members, err = rdb.SMembers(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members of myset after removing member1:")
	for _, member := range members {
		fmt.Println(member)
	}

	// expire a set after 5 seconds
	err = rdb.Expire(ctx, "myset", 5*time.Second).Err()
	if err != nil {
		panic(err)
	}
}
