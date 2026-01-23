package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	err := slowOperation(ctx)
	fmt.Println("Result:", err)
}

func slowOperation(ctx context.Context) error {
	select {
	case <-time.After(1 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

/**
	What this proves:

		Time is enforced outside the operation

		The operation cooperates by listening

		Cancellation is deterministic, not ad hoc
**/
