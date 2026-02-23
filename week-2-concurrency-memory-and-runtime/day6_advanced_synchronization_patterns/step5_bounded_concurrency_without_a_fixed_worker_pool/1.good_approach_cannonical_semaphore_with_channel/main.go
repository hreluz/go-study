package main

import (
	"fmt"
	"sync"
	"time"
)

func work(i int) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("done", i)
}

func main() {
	const limit = 3

	sem := make(chan struct{}, limit) // semaphore tokens
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		sem <- struct{}{} // acquire (blocks when limit reached)
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }() // release
			work(i)
		}(i)
	}

	wg.Wait()
}
