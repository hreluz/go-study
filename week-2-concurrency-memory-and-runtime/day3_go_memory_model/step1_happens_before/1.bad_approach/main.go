package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	var ready bool

	go func() {
		x = 1
		ready = true
	}()

	// Wait a bit (this is NOT synchronization)
	time.Sleep(10 * time.Millisecond)

	if ready {
		fmt.Println("x =", x) // may print 0 or 1; race detector should complain
	} else {
		fmt.Println("not ready")
	}
}

// go run -race main.go
