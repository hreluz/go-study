package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker: context cancelled, returning")
			return
		default:
			fmt.Println("worker: working")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)
	time.Sleep(500 * time.Millisecond)

	cancel() // closes ctx.Done()
	time.Sleep(200 * time.Millisecond)
}

/**
Cause → effect

context.WithCancel creates:

	a context,

	a cancel() function.

	Calling cancel():

	closes the internal Done() channel,

	broadcasts cancellation.

	The goroutine observes <-ctx.Done() and returns.


Accept context.Context as the first parameter of any function that may block or spawn goroutines.
**/
