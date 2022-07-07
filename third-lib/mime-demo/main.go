package main

import (
	"fmt"
	"os"

	"github.com/h2non/filetype"
)

func main() {
	// Open a file descriptor
	file, _ := os.Open("2012a的副本")

	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)

	if filetype.IsDocument(head) {
		fmt.Println("File is an doc")
	} else {
		fmt.Println("Not an doc")
	}
}
