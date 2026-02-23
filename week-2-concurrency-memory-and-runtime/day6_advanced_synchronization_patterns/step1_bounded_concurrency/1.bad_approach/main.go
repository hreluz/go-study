package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func work(i int) {
	time.Sleep(50 * time.Millisecond)
	_ = i * i
}

func main() {
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < 100_000; i++ { // huge fan-out
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			work(i)
		}(i)
	}

	/**
	Each loop iteration:

		1.wg.Add(1) increases the counter.

		2.go func(...) launches a new goroutine immediately.

	There is no limit.

	So after the loop finishes:

		100,000 goroutines have been created.

		They are all either:

			running,

			waiting to run,

			or sleeping.
	**/
	wg.Wait()
	fmt.Println("Num GoRoutine:", runtime.NumGoroutine())
	fmt.Println("done in:", time.Since(start))
}
