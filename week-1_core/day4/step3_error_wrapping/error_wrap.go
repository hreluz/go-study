package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("division by zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func compute() error {
	_, err := divide(10, 0)
	if err != nil {
		// Wrap preserving the cause
		return fmt.Errorf("computing result: %w", err)
	}
	return nil
}

func main() {
	err := compute()
	if err != nil {
		fmt.Println("error:", err)

		// Check root cause
		if errors.Is(err, ErrDivideByZero) {
			fmt.Println("root cause detected: division by zero")
		}
	}
}
