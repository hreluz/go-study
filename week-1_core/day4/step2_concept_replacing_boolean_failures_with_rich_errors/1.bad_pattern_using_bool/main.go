package main

import (
	"fmt"
)

func isValidAge(age int) bool {
	if age < 0 {
		return false
	}
	if age > 120 {
		return false
	}
	return true
}

func main() {
	tests := []int{-5, 30, 200}

	for _, t := range tests {
		ok := isValidAge(t)
		if !ok {
			fmt.Printf("age %d is invalid, but I don't know why\n", t)
			continue
		}
		fmt.Printf("age %d is valid!\n", t)
	}
}
