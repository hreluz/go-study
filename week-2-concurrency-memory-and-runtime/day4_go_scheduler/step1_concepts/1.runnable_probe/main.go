package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
}

/**
	What this proves:

	Your machine has NumCPU hardware cores available to the process.

	The runtime currently allows at most GOMAXPROCS P’s to run Go code simultaneously.

	NumGoroutine shows how many goroutine values exist, not how many are running at once.
**/
