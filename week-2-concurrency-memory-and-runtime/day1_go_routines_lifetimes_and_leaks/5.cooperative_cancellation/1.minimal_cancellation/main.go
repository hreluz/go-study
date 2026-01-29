package main

import (
	"fmt"
	"time"
)

func worker(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("worker: cancelled, returning")
			return
		default:
			fmt.Println("worker: working")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan struct{})

	go worker(done)
	time.Sleep(500 * time.Millisecond)

	close(done) // broadcast cancellation
	time.Sleep(200 * time.Millisecond)
}

/**
Cause → effect

	worker loops forever unless cancellation is observed.

	done is a read-only signal channel.

	Closing the channel wakes all listeners.

	The goroutine chooses to return.


creator goroutine
   |
   +-- owns "done"
   +-- closes "done"
   |
worker goroutine
   |
   +-- observes "done"
   +-- returns
**/
