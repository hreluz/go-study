package main

import (
	"fmt"
	"time"
)

func main() {
	var x int

	go func() {
		x = 42
	}()

	// Timing-based assumption: "it probably finished"
	time.Sleep(10 * time.Millisecond)

	fmt.Println("x =", x) // race: may print 0 or 42; -race should report a data race
}

// go run -race main.go

/**
Why it is wrong (precise)

	There is no happens-before between the write x = 42 and the read fmt.Println(x).

	Sleep does not synchronize memory.

	Therefore, visibility is not guaranteed.
**/
