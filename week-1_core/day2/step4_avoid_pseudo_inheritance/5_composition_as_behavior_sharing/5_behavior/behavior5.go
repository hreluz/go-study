package main

import "fmt"

type AreaCalculator interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 { return s.Side * s.Side }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 { return 3.14159 * c.Radius * c.Radius }

func PrintArea(a AreaCalculator) {
	fmt.Println("Area:", a.Area())
}

func main() {
	PrintArea(Square{Side: 4})
	PrintArea(Circle{Radius: 2})
}

/**
	No hierarchy.
	Completely different structures.
	Only requirement: implement the behavior Area().
**/
