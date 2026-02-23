package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Println("worker", id, "processing", job)
		time.Sleep(200 * time.Millisecond)
		wg.Done()
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, 5) // buffer controls backpressure
	var wg sync.WaitGroup

	// Start fixed workers
	for w := 0; w < numWorkers; w++ {
		go worker(w, jobs, &wg)
	}

	// Submit jobs
	for j := 0; j < numJobs; j++ {
		wg.Add(1)
		jobs <- j
	}

	close(jobs) // important

	wg.Wait()
}
