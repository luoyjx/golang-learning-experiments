package main

import (
	"bytes"
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/mritd/chinaid"
)

func main() {
	bb := bytes.Buffer{}

	for i := 0; i < 1000000; i++ {
		bb.WriteString(chinaid.IDNo())
	}

	fmt.Println(humanize.Bytes(uint64(len(bb.Bytes()))))
}
