package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:0]

	b = append(b, 1)
	b = append(b, 2)
	b = append(b, 3)

	fmt.Println(b, len(b), cap(b))
}
