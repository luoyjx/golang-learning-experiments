package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	simpleConcurrency()
	concurrencyWithBufferedChan()
	concurrencyWithWaitGroup()
}

func simpleConcurrency() {
	done := make(chan int)

	go func() {
		fmt.Println("start job")
		time.Sleep(time.Second)
		fmt.Println("job done")
		<-done
	}()

	done <- 1
	fmt.Println("exit ...")
}

func concurrencyWithBufferedChan() {
	jobs := []int{1, 2, 3, 4, 5}

	startChan := make(chan int, 2)

	for _, j := range jobs {
		go func(v int) {
			fmt.Println("start job", v)
			time.Sleep(time.Second)
			fmt.Println("job done")
			startChan <- 1
		}(j)
	}

	for i := 0; i < len(jobs); i++ {
		<-startChan
	}

	fmt.Println("exit ...")
}

func concurrencyWithWaitGroup() {
	jobs := []int{1, 2, 3, 4, 5}

	wg := sync.WaitGroup{}

	for _, j := range jobs {
		wg.Add(1)

		go func(v int) {
			fmt.Println("start job", v)
			time.Sleep(time.Second)
			fmt.Println("job done")
			wg.Done()
		}(j)
	}

	wg.Wait()

	fmt.Println("exit ...")
}