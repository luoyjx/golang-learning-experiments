package main

import (
	"fmt"
	"time"
)

func main() {
	limitedConcurrency()
}

func limitedConcurrency() {
	works := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	limit := make(chan int, 3)

	for _, work := range works {
		limit <- work
		go func(work int) {
			defer func() {
				<-limit
			}()
			fmt.Println("do work", work)
			time.Sleep(time.Millisecond * 100)
		}(work)
	}

	// if we don't wait here, the program will exit before all the goroutines are done
	time.Sleep(time.Second)
	fmt.Println("works done")
}
