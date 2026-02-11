package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go task("A")
	go task("B")

	time.Sleep(1 * time.Second)
}

/**
	What this proves:

		Two goroutines exist concurrently.

		Output interleaves.

		On a single core, execution is not simultaneous, only interleaved.
**/
