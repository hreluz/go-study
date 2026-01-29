package main

import (
	"context"
	"fmt"
	"time"
)

func workerA(ctx context.Context, errCh chan<- error) {
	// ✅ GOOD: worker does NOT have access to cancel().
	// It can only observe ctx.Done() and report results.
	fmt.Println("A: start; will report an error in 600ms (GOOD)")
	time.Sleep(600 * time.Millisecond)

	// Report “I am done / I failed” upward.
	errCh <- fmt.Errorf("A: encountered an error")
	fmt.Println("A: reported error to owner (GOOD)")
}

func workerB(ctx context.Context) {
	fmt.Println("B: start; should stop only when OWNER cancels")
	t := time.NewTicker(250 * time.Millisecond)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("B: stopped because OWNER cancelled:", ctx.Err())
			return
		case <-t.C:
			fmt.Println("B: working...")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)

	go workerB(ctx)
	go workerA(ctx, errCh)

	// ✅ Owner decides policy: cancel everyone if A fails.
	select {
	case err := <-errCh:
		fmt.Println("main: received from A:", err)
		fmt.Println("main: owner decides to cancel shared ctx")
		cancel()
	case <-time.After(2 * time.Second):
		fmt.Println("main: timeout; owner decides to cancel")
		cancel()
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println("main: exit")
}
