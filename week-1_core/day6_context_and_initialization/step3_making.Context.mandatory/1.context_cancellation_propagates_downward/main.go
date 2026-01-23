package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// The caller defines the operation lifetime
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	err := fetchUser(ctx)
	fmt.Println("Result:", err)
}

func fetchUser(ctx context.Context) error {
	return callDatabase(ctx)
}

func callDatabase(ctx context.Context) error {
	select {
	case <-time.After(1 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// To prove that the caller controls the lifetime, and all callees must obey it.
