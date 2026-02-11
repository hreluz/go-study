package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	const workers = 2
	jobs := make(chan int, 1000)

	var done int32
	var processed int64

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				_ = job // simulate small unit of work
				atomic.AddInt64(&processed, 1)
				if atomic.LoadInt32(&done) == 1 {
					return
				}
			}
		}()
	}

	start := time.Now()
	for time.Since(start) < 2*time.Second {
		jobs <- 1
	}
	atomic.StoreInt32(&done, 1)
	close(jobs)
	wg.Wait()

	fmt.Println("processed:", atomic.LoadInt64(&processed))
}

/**
	Why this is “good”

		You cap concurrency with a fixed number of workers.

		The runtime schedules a small, stable set of goroutines.

		You reduce overhead and reduce imbalance pressure.

		Real-life analogy:

		Instead of creating a new worker for every tiny task, you keep a fixed number of workers who continuously pick tasks from an inbox.
**/
