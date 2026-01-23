package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	go work(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Main finished")
}

func work(ctx context.Context) {
	// Context is ignored here
	time.Sleep(1 * time.Second)
	fmt.Println("Work completed (should not have)")
}
