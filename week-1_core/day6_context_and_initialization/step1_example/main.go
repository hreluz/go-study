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
	fmt.Println("Client left. Canceling...")
	cancel()

	time.Sleep(500 * time.Millisecond)
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

/**
	The caller owns the context because only the caller knows when the operation no longer matters.

	More precisely:

		The caller defines the boundary of meaning of the operation.

		The callee only performs work within that boundary.

		If the callee created or controlled the context, it would be deciding when the caller’s operation ends, which is a violation of responsibility.
**/
