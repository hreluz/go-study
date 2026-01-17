package main

import "fmt"

type Engine struct {
	Horsepower int
}

func (e Engine) Power() int {
	return e.Horsepower
}

type CarB struct {
	Brand  string
	Engine // embedded field
}

func main() {
	b := CarB{
		Brand:  "Tesla",
		Engine: Engine{Horsepower: 350},
	}

	fmt.Println("Brand:", b.Brand)

	// Promoted field:
	fmt.Println("Horsepower:", b.Horsepower)

	// Promoted method:
	fmt.Println("Power():", b.Power())
}
