package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(300 * time.Millisecond)
	fmt.Println("Canceling operation")
	cancel()

	time.Sleep(300 * time.Millisecond)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
