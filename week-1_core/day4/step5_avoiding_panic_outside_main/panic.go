package main

import (
	"errors"
	"fmt"
)

//
// ❌ Version A — Incorrect: using panic for normal errors
//

func loadConfigPanic(path string) string {
	if path == "" {
		panic("config path cannot be empty")
	}
	return "dummy config (panic version)"
}

//
// ✔️ Version B — Correct: returning an error and letting the caller decide
//

func loadConfigError(path string) (string, error) {
	if path == "" {
		return "", errors.New("config path cannot be empty")
	}
	return "dummy config (error version)", nil
}

func main() {

	fmt.Println("---- Demonstrating panic (bad) ----")
	// Recover to avoid process crash
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic recovered in main:", r)
			}
		}()
		_ = loadConfigPanic("") // will panic
	}()

	fmt.Println("\n---- Demonstrating error return (good) ----")
	cfg, err := loadConfigError("") // will return error
	if err != nil {
		fmt.Println("received error instead of panic:", err)
		return
	}
	fmt.Println("config:", cfg)
}
