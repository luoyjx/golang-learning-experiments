package main

import "fmt"

func main() {
	arr := make([]int, 0)
	// will panic
	fmt.Println(arr[0])
}
