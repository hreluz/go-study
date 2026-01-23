package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Manager: STOP")
		cancel()
	}()

	start := time.Now()
	err := packBoxesBad(ctx)
	fmt.Println("packBoxesBad returned after", time.Since(start), "err:", err)
}

func packBoxesBad(ctx context.Context) error {
	for i := 1; i <= 3; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// BAD: "one unit" is huge and not cancelable.
			fmt.Println("BAD: packing box", i, "(takes 1s, cannot be interrupted)")
			time.Sleep(1 * time.Second) // cannot be canceled
		}
	}
	return nil
}
