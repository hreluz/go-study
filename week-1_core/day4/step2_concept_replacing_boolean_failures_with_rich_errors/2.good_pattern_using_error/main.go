package main

import (
	"fmt"
)

func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age %d is negative", age)
	}
	if age > 120 {
		return fmt.Errorf("age %d is unrealistic", age)
	}
	return nil
}

func main() {
	tests := []int{-5, 30, 200}

	for _, t := range tests {
		if err := validateAge(t); err != nil {
			fmt.Printf("validation failed for age %d: %v\n", t, err)
			continue
		}
		fmt.Printf("age %d is valid!\n", t)
	}
}
