package main

import "fmt"

func main() {
	var x int
	ch := make(chan struct{})

	go func() {
		x = 1
		close(ch) // close happens-before the receive unblocks
	}()

	<-ch                  // synchronization point (receive)
	fmt.Println("x =", x) // guaranteed to print 1
}

/**
	Cause → effect: close(ch) synchronizes with <-ch unblocking.
	Therefore, the write x = 1 becomes visible to the reader after <-ch.
**/
