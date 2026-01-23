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

	// GOOD: block waiting for either jobs or cancellation
	go func() {
		start := time.Now()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("GOOD worker exit after %v; err=%v\n", time.Since(start), ctx.Err())
				return
			case job := <-jobs:
				fmt.Println("processing job:", job)
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("Main done:", ctx.Err())
}
