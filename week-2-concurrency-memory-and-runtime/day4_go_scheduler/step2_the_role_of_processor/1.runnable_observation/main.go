package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				fmt.Println("goroutine", id)
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
}

/**
What this demonstrates:

	3 goroutines exist

	1 P exists

	Only one goroutine executes at a time

	Others are runnable but waiting for the P
**/
