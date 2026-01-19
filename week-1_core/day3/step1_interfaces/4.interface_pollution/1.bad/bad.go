package main

import "fmt"

// ---------- INTERFACE POLLUTION EXAMPLE ----------

// Bad interface: only one implementer, same package, no real consumer.
type Calculator interface {
	Compute(a, b int) int
}

// Concrete implementation.
type AddCalculator struct{}

func (AddCalculator) Compute(a, b int) int {
	return a + b
}

// Bad factory: returns the interface, hides the real type.
func NewCalculator() Calculator {
	return AddCalculator{}
}

func main() {
	calc := NewCalculator()

	result := calc.Compute(3, 4)
	fmt.Println("result:", result)

	// PROBLEM:
	// If AddCalculator had additional useful methods (e.g. SetMode),
	// calc could not access them — method set is lost behind the interface.

	// Uncommenting this line would cause a compile error:
	//
	// calc.SetMode("scientific")
	//
	// Because calc is *Calculator*, not *AddCalculator*.
}
