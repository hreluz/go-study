package main

import "fmt"

type Logger struct{}

func (Logger) Log() { fmt.Println("logging") }

type Metrics struct{}

func (Metrics) Log() { fmt.Println("recording metrics") }

type System struct {
	Logger  // behavior 1
	Metrics // behavior 2
}

func main() {
	s := System{}

	// Explicit qualification (Solution A)
	s.Logger.Log()
	s.Metrics.Log()
}
