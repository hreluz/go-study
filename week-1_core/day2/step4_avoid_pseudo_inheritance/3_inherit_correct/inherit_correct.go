package main

import "fmt"

// ❌ Wrong idea: Put behavior in a "base struct"
type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "generic sound"
}

// ❌ Dog is NOT an Animal. This is not OO inheritance.
type DogWrong struct {
	Animal
}

// ✔️ Correct idea: Put common DATA in a reusable struct.
// No inheritance intention.
type Identity struct {
	Name string
}

type Dog struct {
	Identity
	Breed string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	// Wrong design
	wrong := DogWrong{Animal{Name: "Fido"}}
	fmt.Println("Wrong design speaks:", wrong.Speak())

	// Correct design
	good := Dog{
		Identity: Identity{Name: "Rocky"},
		Breed:    "Labrador",
	}
	fmt.Println("Correct design speaks:", good.Speak())
}
