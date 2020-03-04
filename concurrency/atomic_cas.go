package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var G_Int int64
var WG sync.WaitGroup
var ThreadCnt int

// AtomicOpr XXX
func AtomicOpr() {
	var tempInt int64
	for {
		if ThreadCnt == 50 {
			break
		}

		time.Sleep(time.Nanosecond * 5)
	}

	for {
		tempInt = atomic.LoadInt64(&G_Int)
		res := atomic.CompareAndSwapInt64(&G_Int, tempInt, tempInt+1)

		if res == true {
			fmt.Println(tempInt, " try to CAS : ", res)
			break
		}

		time.Sleep(time.Nanosecond * 10)
	}

	WG.Done()
}

func main() {
	G_Int = 0
	ThreadCnt = 0

	for i := 0; i < 50; i++ {
		go AtomicOpr()
		WG.Add(1)

		ThreadCnt++
		fmt.Println("ThreadCnt is :", ThreadCnt)
	}

	WG.Wait()
	time.Sleep(time.Second * 2)
}
