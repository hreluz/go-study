package main

import (
	"runtime"
	"sync"
)

func busy(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	go busy(&wg)
	go busy(&wg)

	wg.Wait()
}

/**
	What this shows:

	Two goroutines

	Two P’s

	Both cores stay busy
**/
