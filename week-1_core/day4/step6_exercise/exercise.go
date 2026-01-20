package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s (%s)", e.Msg, e.Field)
}

func validateAge(age int) error {
	if age < 0 {
		return ValidationError{Field: "age", Msg: "cannot be negative"}
	}
	if age > 120 {
		return ValidationError{Field: "age", Msg: "unrealistic value"}
	}
	return nil
}

func findUser(id int) (string, error) {
	if id != 42 {
		return "", ErrUserNotFound
	}
	return "Héctor", nil
}

func loadProfile(id int, age int) error {
	if err := validateAge(age); err != nil {
		return fmt.Errorf("validating input: %w", err)
	}

	_, err := findUser(id)
	if err != nil {
		return fmt.Errorf("fetching user: %w", err)
	}

	return nil
}

func main() {
	err := loadProfile(7, 200)
	if err == nil {
		fmt.Println("profile loaded")
		return
	}

	fmt.Println("error:", err)

	var ve ValidationError
	if errors.As(err, &ve) {
		fmt.Println("→ validation problem on field:", ve.Field)
		return
	}

	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("→ user does not exist")
	}
}
