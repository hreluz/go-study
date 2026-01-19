package main

import "fmt"

type AddCalculator struct {
	Mode string
}

func (a *AddCalculator) Compute(x, y int) int {
	if a.Mode == "double" {
		return 2 * (x + y)
	}
	return x + y
}

func (a *AddCalculator) SetMode(mode string) {
	a.Mode = mode
}

// Good factory: return concrete.
func NewAddCalculator() *AddCalculator {
	return &AddCalculator{Mode: "normal"}
}

// Accept an interface only at the usage site.
type Computable interface {
	Compute(x, y int) int
}

func UseCalculator(c Computable) {
	fmt.Println("result:", c.Compute(3, 4))
}

func main() {
	calc := NewAddCalculator()

	// You can use it as a concrete type:
	calc.SetMode("double")

	// But you can still pass it where the behavior is needed:
	UseCalculator(calc)
}
