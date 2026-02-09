package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		x = 42
	}()

	wg.Wait()             // establishes happens-before: Done happens-before Wait returns
	fmt.Println("x =", x) // guaranteed 42, no race
}

/**
Cause → effect

	The goroutine calling wg.Done() happens-before wg.Wait() returns.

	After Wait() returns, reads in main are guaranteed to see the goroutine’s writes.
**/
