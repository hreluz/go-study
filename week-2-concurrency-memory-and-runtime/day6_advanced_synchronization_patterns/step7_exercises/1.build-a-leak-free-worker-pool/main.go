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
	fmt.Println("worker", id, "exiting")
}

func main() {
	const workers = 3
	const buffer = 2
	const numJobs = 10

	jobs := make(chan int, buffer)
	var jobsWG sync.WaitGroup

	// Start workers
	for i := 0; i < workers; i++ {
		go worker(i, jobs, &jobsWG)
	}

	// Submit jobs
	for j := 0; j < numJobs; j++ {
		jobsWG.Add(1)
		jobs <- j
	}

	// Close channel (no more jobs)
	close(jobs)

	// Wait for all jobs to complete
	jobsWG.Wait()

	fmt.Println("all jobs done")
}
