package main

import "time"
import "fmt"

func leak() {
	for {
		fmt.Println("running leak")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go leak()
	time.Sleep(1 * time.Second)
}

/**
	leak() has no path to return.

	It is not cancelled.

	It will live for the entire process lifetime.

	This is a goroutine leak by definition.
**/
