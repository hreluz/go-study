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
	err := packBoxesGood(ctx)
	fmt.Println("packBoxesGood returned after", time.Since(start), "err:", err)
}

func packBoxesGood(ctx context.Context) error {
	for i := 1; i <= 10; i++ {
		// One bounded unit of work
		fmt.Println("GOOD: packing small step", i)

		// Make the waiting cancelable:
		if err := sleepOrCancel(ctx, 80*time.Millisecond); err != nil {
			return err
		}
	}
	return nil
}

func sleepOrCancel(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
