package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, jobsWG *sync.WaitGroup, workersWG *sync.WaitGroup) {
	defer workersWG.Done()

	for job := range jobs {
		fmt.Println("worker", id, "processing", job)
		time.Sleep(200 * time.Millisecond)
		jobsWG.Done()
	}

	fmt.Println("worker", id, "exiting")
}

func main() {
	const workers = 3
	const buffer = 2
	const numJobs = 10

	jobs := make(chan int, buffer)

	var jobsWG sync.WaitGroup
	var workersWG sync.WaitGroup

	// Start workers
	for i := 0; i < workers; i++ {
		workersWG.Add(1)
		go worker(i, jobs, &jobsWG, &workersWG)
	}

	// Submit jobs
	for j := 0; j < numJobs; j++ {
		jobsWG.Add(1)
		jobs <- j
	}

	// Signal no more jobs
	close(jobs)

	// Wait for all jobs to complete
	jobsWG.Wait()

	// Wait for all workers to exit
	workersWG.Wait()

	fmt.Println("all jobs done and workers exited")
}

/**
Production complete → close(jobs)
Jobs complete       → jobsWG.Wait()
Workers exit        → workersWG.Wait()
**/
