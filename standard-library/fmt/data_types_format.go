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

	fmt.Println("General ===========================")
	printAsDefaultStructFormat(f)
	printWithStructFieldName(f)
	printGoSyntaxVal(f)
	printGoSyntaxTypeVal(f)
	printLiteralPercentSign(f)
	fmt.Println("Boolean ===========================")
	printBooleanVal()
	fmt.Println("Integer ===========================")
	printIntBase2Val()
	printIntCharUnicodeCodePoint()
	printIntBase10()
	printIntBase8()
	printIntBase8WithPrefix()
	printSingleQuotedChar()
	printIntBase16LowerCase()
	printIntBase16UpperCase()
	printIntUnicodeFormat()
	fmt.Println("String and slice of bytes ===========================")
	printStringNormal()
	printStringDoubleQuoted()
	printStringBase16Lower()
	printStringBase16Upper()
	fmt.Println("Slice ============================")
	printSlice()
	fmt.Println("Pointer ===========================")
	printPointer()
	fmt.Println("Float ============================")
	printDecimallessScientificNotation()
	printScientificNotationLower()
	printScientificNotationUpper()
	printDecimalPoint()
	printHexadecimalNotation()
	printUpperHexadecimalNotation()
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

func printGoSyntaxVal(f Foo) {
	// print:
	// go syntax: main.Foo{Bar:"barvalue"}
	fmt.Printf("go syntax val: %#v\n", f)
}

func printGoSyntaxTypeVal(f Foo) {
	// print:
	// go syntax type: main.Foo
	fmt.Printf("go syntax type: %T\n", f)
}

func printLiteralPercentSign(f Foo) {
	// print:
	// literal percent sign: %v {barvalue}
	fmt.Printf("literal percent sign: %%v %v \n", f)
}

func printBooleanVal() {
	// print:
	// boolean: true
	fmt.Printf("boolean: %t\n", true)

	// print:
	// boolean: false
	fmt.Printf("boolean: %t\n", false)
}

func printIntBase2Val() {
	// print:
	// int base 2: 100
	fmt.Printf("int base 2: %b\n", 4)
}

func printIntCharUnicodeCodePoint() {
	// print:
	// int char unicode code point: a
	fmt.Printf("int char unicode code point: %c\n", 97)

	// print:
	// int char unicode code point: a
	fmt.Printf("int char unicode code point: %c\n", 'a')
}

func printIntBase10() {
	// print:
	// int base 10: 100
	fmt.Printf("int base 10: %d\n", 100)
}

func printIntBase8() {
	// print:
	// int base 8: 144
	fmt.Printf("int base 8: %o\n", 100)

	// print:
	// int base 8: 0173
	fmt.Printf("int base 8: %#o\n", 123)
}

func printIntBase8WithPrefix() {
	// print:
	// int base 8 with prefix: 0o173
	fmt.Printf("int base 8 with prefix: %O\n", 123)
}

func printSingleQuotedChar() {
	// print:
	// single quoted char: 'a'
	fmt.Printf("single quoted char: %q\n", 97)
}

func printIntBase16LowerCase() {
	// print:
	// int base 16 lower case: 3e8
	fmt.Printf("int base 16 lower case: %x\n", 1000)
}

func printIntBase16UpperCase() {
	// print:
	// int base 16 upper case: 3E8
	fmt.Printf("int base 16 upper case: %X\n", 1000)
}

func printIntUnicodeFormat() {
	// print:
	// int unicode format: U+0064
	fmt.Printf("int unicode format: %U\n", 100)
}

func printStringNormal() {
	// print:
	// string normal: abc
	fmt.Printf("string normal: %s\n", "abc")

	// print:
	// string normal: abc
	fmt.Printf("string normal: %s\n", []byte("abc"))
}

func printStringDoubleQuoted() {
	// print:
	// string double quoted: "abc"
	fmt.Printf("string double quoted: %q\n", "abc")
}

func printStringBase16Lower() {
	// print:
	// string base 16 lower: 616263
	fmt.Printf("string base 16 lower: %x\n", "abc")
}

func printStringBase16Upper() {
	// print:
	// string base 16 upper: 616263
	fmt.Printf("string base 16 upper: %X\n", "abc")
}

func printSlice() {
	// print some thing like:
	// slice: 0xc0000cc060
	fmt.Printf("slice: %p\n", []int{1, 2, 3})
}

func printPointer() {
	// print something like:
	// pointer: 0xc0000cc060
	fmt.Printf("pointer: %p\n", &[]int{1, 2, 3})
}

func printDecimallessScientificNotation() {
	// print:
	// decimalless scientific notation: -1111011
	fmt.Printf("decimalless scientific notation: %b\n", -123)
}

func printScientificNotationLower() {
	// print:
	// scientific notation lower: -1.234568e+05
	fmt.Printf("scientific notation lower: %e\n", -123456.789)
}

func printScientificNotationUpper() {
	// print:
	// scientific notation upper: -1.234568E+05
	fmt.Printf("scientific notation upper: %E\n", -123456.789)
}

func printDecimalPoint() {
	// print:
	// decimal point: 123.456000
	fmt.Printf("decimal point: %f\n", 123.456)
}

func printHexadecimalNotation() {
	// print:
	// hexadecimal notation: 186a0
	fmt.Printf("hexadecimal notation: %x\n", 100000)
}

func printUpperHexadecimalNotation() {
	// print:
	// upper hexadecimal notation: 186A0
	fmt.Printf("upper hexadecimal notation: %X\n", 100000)
}
