package main

// Import the redis package
import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Add an item to the list
	err := rdb.LPush(context.Background(), "mylist", "item1").Err()
	if err != nil {
		panic(err)
	}

	// Add multiple items to the list
	err = rdb.LPush(context.Background(), "mylist", "item2", "item3").Err()
	if err != nil {
		panic(err)
	}

	// Get the length of the list
	length, err := rdb.LLen(context.Background(), "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of list:", length)

	// Get the first item in the list
	firstItem, err := rdb.LIndex(context.Background(), "mylist", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("First item in list:", firstItem)

	// Get all items in the list
	allItems, err := rdb.LRange(context.Background(), "mylist", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All items in list:", allItems)

	// Remove the last item from the list
	err = rdb.RPop(context.Background(), "mylist").Err()
	if err != nil {
		panic(err)
	}
}
