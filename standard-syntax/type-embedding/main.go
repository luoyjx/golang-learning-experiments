package main

import "fmt"

type Lib interface {
	Put()
}

type libImpl struct{}

func (l *libImpl) Put() {
	fmt.Println("put")
}

type Base struct {
	LibClient Lib
}

func (b *Base) PutSomething() {
	b.LibClient.Put()
}

func (b *Base) GetSomething() {
	fmt.Println("base get")
}

type Foo struct {
	Base
}

func (b *Foo) PutSomething() {
	fmt.Println("foo put")
}

type Bar struct {
	*Base
}

func (b *Bar) PutSomething() {
	fmt.Println("bar put")
}

func main() {
	lib := &libImpl{}
	foo := &Foo{
		Base: Base{LibClient: lib},
	}

	foo.PutSomething()
	foo.GetSomething()

	bar := &Bar{
		Base: &Base{LibClient: lib},
	}

	bar.GetSomething()
	bar.PutSomething()
}
