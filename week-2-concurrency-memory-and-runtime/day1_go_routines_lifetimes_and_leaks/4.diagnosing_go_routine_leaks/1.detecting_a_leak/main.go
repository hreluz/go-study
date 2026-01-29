package main

import (
	"fmt"
	"runtime"
	"time"
)

func leak() {
	select {} // blocked forever
}

func main() {
	fmt.Println("initial:", runtime.NumGoroutine())

	go leak()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("after leak:", runtime.NumGoroutine())
}

/**
Cause → effect
	leak never returns.

	Goroutine count increases.

	The count never goes back down.

	This is objective evidence of a leak.
**/
