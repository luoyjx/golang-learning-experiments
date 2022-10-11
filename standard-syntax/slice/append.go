package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["a"] = append(m["a"], "b")
	fmt.Println("append to nil slice:", m["a"])
}
