package main

import (
	"bytes"
	"fmt"
)

// Greeter defines a behavioral contract.
// Any type that has Greet(string) string satisfies this interface.
type Greeter interface {
	Greet(name string) string
}

// EnglishGreeter is a concrete type.
type EnglishGreeter struct{}

// Greet implements the Greeter behavior.
func (EnglishGreeter) Greet(name string) string {
	return "Hello, " + name
}

// UppercaseGreeter wraps another Greeter and transforms the output.
type UppercaseGreeter struct {
	Base Greeter
}

// Greet implements the Greeter interface.
func (u UppercaseGreeter) Greet(name string) string {
	original := u.Base.Greet(name)
	return string(bytes.ToUpper([]byte(original)))
}

// PrintGreeting depends only on the behavior, not on concrete types.
func PrintGreeting(g Greeter, name string) {
	fmt.Println(g.Greet(name))
}

func main() {
	english := EnglishGreeter{}
	upper := UppercaseGreeter{Base: english}

	PrintGreeting(english, "Héctor")
	PrintGreeting(upper, "Héctor")
}
