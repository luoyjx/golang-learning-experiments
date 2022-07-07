package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := arr1[1:]
	fmt.Println(arr1)
	arr2 = append(arr2, 4)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
