package main

import (
	"context"
	"fmt"
	"time"
)

func workerA(ctx context.Context) {
	// ❌ BAD: worker cancels the shared context it did not create.
	// This terminates other goroutines that depend on the same ctx.
	fmt.Println("A: start; will cancel shared ctx in 600ms (BAD)")
	time.Sleep(600 * time.Millisecond)

	// Pretend A decided it is "done" and cancels everything.
	// In real systems, this might happen via defer cancel() or an error path.
	cancelShared(ctx) // shows the design flaw
	fmt.Println("A: called cancel on shared ctx (BAD)")
}

func workerB(ctx context.Context) {
	fmt.Println("B: start; expects to run ~2s unless OWNER cancels")
	t := time.NewTicker(250 * time.Millisecond)
	defer t.Stop()

	deadline := time.NewTimer(2 * time.Second)
	defer deadline.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("B: stopped EARLY due to shared cancellation:", ctx.Err())
			return
		case <-t.C:
			fmt.Println("B: working...")
		case <-deadline.C:
			fmt.Println("B: finished normally (this should happen if not cancelled)")
			return
		}
	}
}

// We cannot extract cancel() from ctx directly, so we simulate the bad pattern
// by storing the cancel func in a place the worker can access.
// This is exactly the kind of “ownership leak” that happens in real code
// when cancel is passed down, stored on a struct, or put in a global.
var sharedCancel func()

func cancelShared(ctx context.Context) {
	if sharedCancel != nil {
		sharedCancel()
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sharedCancel = cancel // ❌ BAD: exposes cancel to non-owners

	go workerB(ctx)
	go workerA(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("main: exit")
}

/**
	workerB expects to run ~2 seconds.

	workerA cancels the shared context at ~600ms.

	workerB stops early, even though the owner (main) did not decide to stop it.
**/
