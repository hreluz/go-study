package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	fmt.Println("Before declaring go routine")

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- "done"
	}()

	fmt.Println("After declaring go routine")

	select {
	case msg := <-ch:
		fmt.Println("received:", msg)
	}

	fmt.Println("After select ...")
}
