package main

import "fmt"

// "Base" type with both data and behavior
type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "generic sound"
}

// ❌ Inheritance-style thinking: Dog "is an" Animal
type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func main() {
	d := Dog{Animal{Name: "Fido"}}
	c := Cat{Animal{Name: "Misu"}}

	fmt.Println("Dog name:", d.Name)
	fmt.Println("Dog speaks:", d.Speak()) // generic sound

	fmt.Println("Cat name:", c.Name)
	fmt.Println("Cat speaks:", c.Speak()) // generic sound
}
