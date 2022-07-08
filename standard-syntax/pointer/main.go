package main

import "fmt"

type A struct {
	Name   string
	Age    int
	Parent *A
}

func main() {
	a := new(A)
	fmt.Println(a)
}
