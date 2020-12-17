package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Foo string
}

type B struct {
	Foo string
	Bar string
}

func main() {
	b := &B{"a", "b"}

	bs, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}

	var a *A

	err = json.Unmarshal(bs, a)

	if err != nil {
		panic(err)
	}

	fmt.Println(a.Foo)

}
