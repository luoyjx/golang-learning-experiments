
// Define the interface Lib
type Lib interface {
	Put()
}

// Define the struct libImpl that implements the Lib interface
type libImpl struct{}

// Implement the Put method of the Lib interface for the libImpl struct
func (l *libImpl) Put() {
	fmt.Println("put")
}

// Define the struct Base that has a LibClient field of type Lib interface
type Base struct {
	LibClient Lib
}

// Implement the PutSomething method for the Base struct that calls the Put method of the LibClient field
func (b *Base) PutSomething() {
	b.LibClient.Put()
}

// Implement the GetSomething method for the Base struct
func (b *Base) GetSomething() {
	fmt.Println("base get")
}

// Define the struct Foo that embeds the Base struct
type Foo struct {
	Base
}

// Override the PutSomething method for the Foo struct
func (b *Foo) PutSomething() {
	fmt.Println("foo put")
}

// Define the struct Bar that embeds a pointer to the Base struct
type Bar struct {
	*Base
}

// Override the PutSomething method for the Bar struct
func (b *Bar) PutSomething() {
	fmt.Println("bar put")
}

func main() {
	// Create an instance of the libImpl struct
	lib := &libImpl{}

	// Create an instance of the Foo struct that has a Base field with a LibClient field set to the libImpl instance
	foo := &Foo{
		Base: Base{LibClient: lib},
	}

	// Call the PutSomething method of the Foo struct
	foo.PutSomething()

	// Call the GetSomething method of the Foo struct
	foo.GetSomething()

	// Create an instance of the Bar struct that has a Base field with a LibClient field set to the libImpl instance
	bar := &Bar{
		Base: &Base{LibClient: lib},
	}

	// Call the GetSomething method of the Bar struct
	bar.GetSomething()

	// Call the PutSomething method of the Bar struct
	bar.PutSomething()
}

