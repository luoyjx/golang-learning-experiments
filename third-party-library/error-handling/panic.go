package main

import (
	"fmt"

	"github.com/rotisserie/eris"
)

func funcB() error {
	return eris.New("some error")
}

func funcA() error {
	if err := funcB(); err != nil {
		return eris.Wrap(err, "funcB error")
	}

	return nil
}

func main() {
	err := funcA()
	if err != nil {
		panic(err)
	}

	fmt.Println("ok fine")
}
