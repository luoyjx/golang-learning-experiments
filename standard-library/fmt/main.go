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
	fmt.Println("print Foo struct JSON bytes as string")
	fmt.Printf("print Foo JSON bytes string with fmt.Printf %s\n", string(bs))
	fmt.Println("fmt print format reference => https://pkg.go.dev/fmt")

	printAsDefaultStructFormat(f)
	printWithStructFieldName(f)
}

func printAsDefaultStructFormat(f Foo) {
	// print:
	// default struct format: {barvalue}
	fmt.Printf("default struct format: %v\n", f)
}

func printWithStructFieldName(f Foo) {
	// print:
	// struct field name: {Bar:barvalue}
	fmt.Printf("struct field name: %+v\n", f)
}
