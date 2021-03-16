package main

import (
	"fmt"
	"strings"
)

func main() {
	s := make([]byte, 4096)
	fmt.Println("s", len(s), cap(s))
	sr := strings.NewReader("hahahahhaha")
	_, _ = sr.Read(s)
	fmt.Println("s", len(s), cap(s))
	s1 := make([]byte, 0, 4096)
	fmt.Println("s1", len(s1), cap(s1))
	sr1 := strings.NewReader("hahahahhaha")
	_, _ = sr1.Read(s)
	fmt.Println("s1", len(s1), cap(s1))
}