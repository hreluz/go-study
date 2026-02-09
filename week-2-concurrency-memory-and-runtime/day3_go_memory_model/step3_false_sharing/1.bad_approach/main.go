package main

import (
	"sync"
)

type Counters struct {
	a int64
	b int64
}

func main() {
	var c Counters
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000_000; i++ {
			c.a++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000_000; i++ {
			c.b++
		}
	}()

	wg.Wait()
}

/**
- Correct? ✔️ (no data race)
- Fast? ❌ (likely slow due to false sharing)

Why?
a and b are likely in the same cache line.
**/
