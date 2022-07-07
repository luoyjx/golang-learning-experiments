package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func ReadFull() {
	s := "some io.Reader stream to be read\n"
	r := strings.NewReader(s)

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err, "len of string ", len(s))
	}
}

func main() {
	fmt.Println("io.ReadFull")
	ReadFull()
}
