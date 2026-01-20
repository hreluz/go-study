package main

import (
	"errors"
	"fmt"
)

// Sentinel error — a shared, package-level value
var ErrUserNotFound = errors.New("user not found")

func findUser(id int) (string, error) {
	if id != 42 {
		// Return the sentinel error directly
		return "", ErrUserNotFound
	}
	return "Héctor", nil
}

func loadProfile(id int) error {
	_, err := findUser(id)
	if err != nil {
		// Wrap the sentinel error to add context
		return fmt.Errorf("loading profile for id=%d: %w", id, err)
	}
	return nil
}

func main() {
	err := loadProfile(7)
	if err == nil {
		fmt.Println("profile loaded")
		return
	}

	fmt.Println("error:", err)

	// Detect sentinel error even after wrapping
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("→ root cause: user does not exist")
	}
}
