package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	// Consumer
	go func() {
		for v := range ch {
			fmt.Println("Processing:", v)
			time.Sleep(1 * time.Second)
		}
	}()

	// Producer
	for i := 0; i < 10; i++ {
		fmt.Println("Producing:", i)
		ch <- i
		fmt.Println("Produced:", i)
	}

	close(ch)
}

/**
After 3 items:

	Producer blocks until consumer processes one.

	You will see pauses in output.
**/

/**
	Why does the producer block after sending 3 items?
		The buffered channel has reached capacity, so the send operation cannot proceed until a receive removes a value from the buffer.
**/
