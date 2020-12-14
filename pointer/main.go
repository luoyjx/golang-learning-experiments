package main

type A struct {
	Name string
	Age int
	Parent *A
}

func main() {
	a := new(A)

	a->Parent = new(A)
	fmt.Println(a)
}