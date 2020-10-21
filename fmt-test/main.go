package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Bar string
}

func main() {
	f := Foo{"barvalue"}

	bs, err := json.Marshal(f)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}