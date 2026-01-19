package main

import "fmt"

// Behavior expected
type Describer interface {
	Describe() string
}

// Concrete type with two methods
type Info struct {
	Value string
}

// Value receiver method
func (i Info) Describe() string {
	return "Info: " + i.Value
}

// Pointer receiver method
func (i *Info) Reset() {
	i.Value = ""
}

// New struct embedding Info
type Item struct {
	Info // embedded directly (value)
}

func main() {
	obj := Item{
		Info: Info{Value: "embedded"},
	}

	// 1. Does Item satisfy Describer?
	fmt.Println(obj.Describe()) // promoted method

	// 2. Can Item call Reset()?
	// The following WOULD NOT COMPILE:
	//
	// obj.Reset()
	//
	// Because obj.Info is a VALUE, so its method set
	// does NOT include pointer-receiver methods.

	// But a pointer to Item works:
	ptr := &obj
	ptr.Reset() // now allowed (Info's pointer methods promoted)

	fmt.Println("After Reset:", obj.Describe())
}
