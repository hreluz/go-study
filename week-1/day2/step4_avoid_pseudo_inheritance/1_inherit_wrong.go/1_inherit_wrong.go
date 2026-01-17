package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "generic sound"
}

type Dog struct {
	Animal // ❌ This feels like inheritance, but is NOT
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Fido"},
	}

	// Many beginners think this becomes "Woof!"
	fmt.Println("Dog speaks:", d.Speak()) // Actually: "generic sound"
}
