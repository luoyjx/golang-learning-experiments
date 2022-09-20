package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// size of 4 chars string
	fmt.Println(unsafe.Sizeof("jack"))
	// size of 44 chars string
	fmt.Println(unsafe.Sizeof("jackjackjackjackjackjackjackjackjackjackjack"))
	// size of 44 chars byte array
	fmt.Println(unsafe.Sizeof([44]byte{'j', 'a', 'a', 'k'}))
}
