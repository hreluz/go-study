package main

import "fmt"

type Engine struct {
	Horsepower int
}

func (e Engine) Power() int {
	return e.Horsepower
}

type CarA struct {
	Brand  string
	Engine Engine // named field
}

func main() {
	a := CarA{
		Brand:  "Toyota",
		Engine: Engine{Horsepower: 150},
	}

	fmt.Println("Brand:", a.Brand)
	fmt.Println("Horsepower:", a.Engine.Horsepower)
	fmt.Println("Power():", a.Engine.Power())
}
