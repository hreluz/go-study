package main

import (
	"errors"
	"fmt"
)

type AgeError struct {
	Age int
	Msg string
}

func (e AgeError) Error() string {
	return fmt.Sprintf("age error: %s (age=%d)", e.Msg, e.Age)
}

func validateAge(age int) error {
	if age < 0 {
		return AgeError{Age: age, Msg: "age cannot be negative"}
	}
	if age > 120 {
		return AgeError{Age: age, Msg: "age unrealistic"}
	}
	return nil
}

func process(age int) error {
	if err := validateAge(age); err != nil {
		return fmt.Errorf("processing user age: %w", err)
	}
	return nil
}

func main() {
	err := process(200)
	if err == nil {
		fmt.Println("all good")
		return
	}

	fmt.Println("error:", err)

	var ae AgeError
	if errors.As(err, &ae) {
		fmt.Println("→ extracted structured info:")
		fmt.Println("   failing age:", ae.Age)
		fmt.Println("   reason:", ae.Msg)
	}
}
