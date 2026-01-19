package main

import "fmt"

// Greeter is a behavioral contract.
type Greeter interface {
	Greet(name string) string
}

// FancyGreeter is a concrete type with extra capabilities.
type FancyGreeter struct {
	Prefix string
}

func (f *FancyGreeter) Greet(name string) string {
	return f.Prefix + " " + name
}

// Extra behavior that is NOT part of the interface.
func (f *FancyGreeter) SetPrefix(p string) {
	f.Prefix = p
}

// -------- BAD DESIGN: constructor returns an interface --------

// NewGreeterBad returns the interface type.
// This hides the concrete type behind the interface.
func NewGreeterBad() Greeter {
	return &FancyGreeter{Prefix: "Hello"}
}

// -------- GOOD DESIGN: constructor returns concrete type --------

// NewGreeterGood returns the concrete type (pointer).
func NewGreeterGood() *FancyGreeter {
	return &FancyGreeter{Prefix: "Hello"}
}

// UseGreeter ACCEPTS an interface.
// It only cares about the behavior.
func UseGreeter(g Greeter, name string) {
	fmt.Println("UseGreeter:", g.Greet(name))
}

func main() {
	fmt.Println("=== Using bad constructor (returns interface) ===")
	bad := NewGreeterBad()
	UseGreeter(bad, "Héctor")

	// The following line WOULD NOT COMPILE if uncommented:
	//
	// bad.SetPrefix("Hi")
	//
	// Because 'bad' is of type Greeter, and Greeter does not declare SetPrefix.
	// So we CANNOT use the extra capabilities of FancyGreeter here.

	fmt.Println("\n=== Using good constructor (returns concrete) ===")
	good := NewGreeterGood()
	UseGreeter(good, "Héctor")

	// Here we still can pass it where a Greeter is needed...
	UseGreeter(good, "Héctor again")

	// ...but we ALSO can use the extra behavior of the concrete type.
	good.SetPrefix("Hi")
	UseGreeter(good, "Héctor after prefix change")
}
