package main

import "time"

func pollForever() {
	for {
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go pollForever()
	select {}
}

/**
Why this leaks

	Still an infinite loop.

	time.Sleep does not create an exit condition.

	The goroutine will never return.

Observable signal

	Low CPU.

	Memory and goroutine count remain elevated.

	Often missed in reviews.
**/
