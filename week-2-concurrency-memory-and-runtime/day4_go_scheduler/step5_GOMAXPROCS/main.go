package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func cpuWork(d time.Duration) {
	end := time.Now().Add(d)
	x := 0
	for time.Now().Before(end) {
		x++
	}
	_ = x
}

func main() {
	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2)

	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { defer wg.Done(); cpuWork(1 * time.Second) }()
	go func() { defer wg.Done(); cpuWork(1 * time.Second) }()

	wg.Wait()
	fmt.Println("elapsed:", time.Since(start))
}

/**
	Expected observation:

		With 1: elapsed close to ~2 seconds (two CPU-bound tasks share one P)

		With 2: elapsed close to ~1 second (two P’s allow parallel execution)

	Note:

		Exact times vary due to CPU frequency scaling and OS noise.

		The trend is what matters.
**/
