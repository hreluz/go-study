package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "generic sound"
}

type Dog struct {
	Animal
}

// This does NOT override Animal.Speak.
// It simply defines a Dog.Speak.
// The promoted method (Animal.Speak) still exists.
func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Fido"},
	}

	// This now resolves to Dog.Speak (method shadowing)
	fmt.Println("Dog speaks:", d.Speak())

	// You can still reach the embedded method explicitly
	fmt.Println("Animal speaks:", d.Animal.Speak())
}
