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
		wg.Done() // decrement unfinished jobs
	}
	fmt.Println("worker", id, "exiting")
}

func main() {
	const workers = 3
	const numJobs = 10

	jobs := make(chan int, 2) // small buffer for backpressure
	var jobsWG sync.WaitGroup

	// Start workers
	for i := 0; i < workers; i++ {
		go worker(i, jobs, &jobsWG)
	}

	// Submit jobs
	for j := 0; j < numJobs; j++ {
		jobsWG.Add(1) // increment unfinished jobs
		jobs <- j
	}

	// Signal no more jobs
	fmt.Println("close(jobs)")
	close(jobs)

	// Wait for all jobs to finish
	fmt.Println("jobsWG.Wait()")
	jobsWG.Wait() // unblocks

	fmt.Println("all jobs done")
}

/**
Why this is leak-free

	Step-by-step:

	1.Workers are started once.

	2.Each job increments jobsWG.

	3.Workers call wg.Done() after finishing each job.

	4.After submitting all jobs → close(jobs).

	5.Workers eventually:

		-Drain remaining jobs.

		-Exit range jobs.

		-Terminate.

	6.jobsWG.Wait() ensures all work is done.

	No worker blocks forever.

	No goroutine leaks.
**/
