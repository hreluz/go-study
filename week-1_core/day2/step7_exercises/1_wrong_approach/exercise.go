package main

import "fmt"

// ------------------------------
// "Inheritance-style" design
// ------------------------------

// Base type mixing data + behavior
type Animal struct {
	Name string
}

func (a Animal) Describe() string {
	return "Animal: " + a.Name
}

// Child types that "inherit" from Animal via embedding.
// This is the inheritance-style mindset we want to avoid.
type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

// ------------------------------
// Second problem: conflicting behaviors
// ------------------------------

type Logger struct{}

func (Logger) Info(msg string) {
	fmt.Println("[INFO]", msg)
}

type AuditLogger struct{}

func (AuditLogger) Info(msg string) {
	fmt.Println("[AUDIT]", msg)
}

// Bad: two embedded types exposing the same method name.
// This causes ambiguity if you try to call ps.Info().
type PaymentService struct {
	Logger
	AuditLogger
}

// ------------------------------
// main (current behavior)
// ------------------------------

func main() {
	// Inheritance-style usage
	dog := Dog{Animal{Name: "Fido"}}
	cat := Cat{Animal{Name: "Misu"}}

	fmt.Println(dog.Describe())
	fmt.Println(cat.Describe())

	// Uncomment this to see the ambiguity compile error:
	ps := PaymentService{}
	ps.Info("payment created") // ambiguous selector
	_ = PaymentService{}
}
