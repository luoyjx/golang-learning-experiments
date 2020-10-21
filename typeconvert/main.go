package main

import (
	"fmt"
	"strconv"
)

func main() {
	uidStr := "-123"
	uid, err := strconv.ParseInt(uidStr, 10, 64)

	if err != nil {
		panic(err)
	}

	fmt.Println(uid)
}
