package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 2, 3}

	// 非缓冲 channel ，如果未被取出时，会被阻塞
	inChan := make(chan struct{})
	doneC := make(chan struct{})

	go func() {
		for _, v := range arr {
			fmt.Printf("in %d \n", v)

			inChan <- struct{}{}
		}

		// 如果不关闭 channel 的话，会引起接收部分的 inChan 一直等待，
		// doneC 也一直等待，死锁
		close(inChan)
	}()

	go func() {
		for range inChan {
			fmt.Println("receve a chan item")
			time.Sleep(time.Duration(time.Second * 5))
		}

		doneC <- struct{}{}
	}()

	<-doneC
}
