package main

import (
	"fmt"
	"sync"
)

func main() {
	increments := make(chan int)
	done := make(chan struct{})

	// Owner goroutine: the ONLY goroutine that touches `count`.
	go func() {
		count := 0
		for inc := range increments { // loop ends when increments is closed
			count += inc
		}
		fmt.Println("final:", count)
		close(done)
	}()

	const workers = 5
	const perWorker = 1000

	var wg sync.WaitGroup
	wg.Add(workers)

	// Senders
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < perWorker; j++ {
				increments <- 1
			}
		}()
	}

	// Close only after all senders are done.
	go func() {
		wg.Wait()
		close(increments)
	}()

	<-done
}
