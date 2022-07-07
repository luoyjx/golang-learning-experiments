package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var gInt int64
var wg sync.WaitGroup
var threadCnt int

// AtomicOpr XXX
func AtomicOpr() {
	var tempInt int64
	for {
		if threadCnt == 50 {
			break
		}

		time.Sleep(time.Nanosecond * 5)
	}

	for {
		tempInt = atomic.LoadInt64(&gInt)
		res := atomic.CompareAndSwapInt64(&gInt, tempInt, tempInt+1)

		if res == true {
			fmt.Println(tempInt, " try to CAS : ", res)
			break
		}

		time.Sleep(time.Nanosecond * 10)
	}

	wg.Done()
}

func main() {
	gInt = 0
	threadCnt = 0

	for i := 0; i < 50; i++ {
		go AtomicOpr()
		wg.Add(1)

		threadCnt++
		fmt.Println("threadCnt is :", threadCnt)
	}

	wg.Wait()
	time.Sleep(time.Second * 2)
}
