package main

import (
	"fmt"
	"sync"
	"time"
)

func work(i int) {
	time.Sleep(50 * time.Millisecond)
	_ = i * i
}

func main() {
	const maxInFlight = 100

	sem := make(chan struct{}, maxInFlight) // capacity = limit
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < 100_000; i++ {
		sem <- struct{}{} // acquire token (blocks when full)
		/**
		It tries to put one empty value into the channel.

			If the channel buffer is full:
			→ it blocks.

		Blocking means:
			the current goroutine (main loop) stops executing until space is available.
		**/
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }() // release token
			fmt.Println(i)
			work(i)
		}(i)
	}

	wg.Wait()
	fmt.Println("done in:", time.Since(start))
}

/**
Visual comparison
	Bad version
		Loop
		├─ spawn
		├─ spawn
		├─ spawn
		├─ spawn
		├─ spawn
		├─ spawn
		...
		(no limit)


	Concurrency grows with loop size.

	Good version
		Loop
		├─ acquire token (blocks if full)
		├─ spawn
		├─ acquire token
		├─ spawn
		...

	If 100 are already running:

		sem buffer full
		↓
		main goroutine blocks
		↓
		no new goroutines created


	Concurrency is capped.
**/

/**
Cause → Effect chain

	1.Channel capacity = N

	2.N workers started

	3.Channel buffer becomes full

	4.Next sem <- struct{}{} blocks

	5.No more goroutines are created

	6.A worker finishes and executes <-sem

	7.Buffer has space

	8.Sender unblocks

	9.Loop continues

	This is mechanical.

	No magic.
**/
