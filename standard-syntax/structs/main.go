package main

import "fmt"

type TestS struct {
	delayMs int64
}

type SomeType struct {
	Version int64
}

func main() {
	t := TestS{}

	fmt.Println("delay ", t.delayMs)

	st := SomeType{}

	fmt.Println("some type version ", int(st.Version))
}
