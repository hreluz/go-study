package main

import (
	"fmt"
	"time"
)

func main() {

	var x int

	go func() {
		x = 3
	}()

	time.Sleep(10 * time.Millisecond)
	fmt.Println(x)
}

/**
	Write a program with:

	One goroutine writing to a variable

	One goroutine reading it

	No synchronization

	Run it with -race.

	Goal: Prove that execution timing does not matter.
**/

//  go run -race main.go
