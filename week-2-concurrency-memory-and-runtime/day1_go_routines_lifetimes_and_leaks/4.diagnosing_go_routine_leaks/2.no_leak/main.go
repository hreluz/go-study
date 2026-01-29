package main

import (
	"fmt"
	"runtime"
	"time"
)

func finite() {
	time.Sleep(100 * time.Millisecond)
}

func main() {
	fmt.Println("initial:", runtime.NumGoroutine())

	go finite()
	time.Sleep(300 * time.Millisecond)

	fmt.Println("after finite:", runtime.NumGoroutine())
}

/**
Cause → effect

	Goroutine starts.

	Goroutine returns.

	Goroutine count returns to baseline.

	This is a bounded lifetime.



Critical limitation (must be internalized)

	runtime.NumGoroutine() cannot tell you:

	which goroutine leaked,

	why it leaked,

	where it was created.

It only answers:

	“Did something start and fail to end?”
**/
