package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	jobs := make(chan int) // no jobs will arrive

	// BAD: a worker that polls jobs with default -> busy loop
	go func() {
		var spins int
		start := time.Now()

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("BAD worker exit after %v; spins=%d; err=%v\n", time.Since(start), spins, ctx.Err())
				return
			default:
				// Nothing to do, but we keep looping immediately.
				// This can burn CPU.
				spins++
			}
		}
	}()

	// Wait for timeout; no jobs ever sent.
	<-ctx.Done()
	fmt.Println("Main done:", ctx.Err())
	_ = jobs
}
