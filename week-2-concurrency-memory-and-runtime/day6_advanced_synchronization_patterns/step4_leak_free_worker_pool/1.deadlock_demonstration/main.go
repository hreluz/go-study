package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	var jobsWG sync.WaitGroup
	var workersWG sync.WaitGroup

	// 1 worker
	workersWG.Add(1)
	go func() {
		defer workersWG.Done()
		for j := range jobs {
			_ = j
			jobsWG.Done()
		}
	}()

	// produce 1 job
	jobsWG.Add(1)
	jobs <- 1

	// WRONG order
	jobsWG.Wait()    // waits until job done (ok)
	workersWG.Wait() // waits for worker to exit (blocked!)
	close(jobs)      // never reached

	fmt.Println("unreachable")
}

/**
	Here it never finishes because:

	worker will not exit until jobs is closed,

	but close(jobs) is after workersWG.Wait().

	That is the “big deal.”
**/
