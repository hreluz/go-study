package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				fmt.Println("goroutine", id, "working")
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

	time.Sleep(5 * time.Second)
}

/**
	What this proves:

	time.Sleep is a blocking syscall

	Goroutines continue executing despite others sleeping

	P is being reused by different M’s
**/
