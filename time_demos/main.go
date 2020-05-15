package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	time.Sleep(time.Duration(500 * time.Millisecond))
	cost := time.Now().Sub(startTime)

	fmt.Println("cost duration ", cost.Milliseconds())
}
