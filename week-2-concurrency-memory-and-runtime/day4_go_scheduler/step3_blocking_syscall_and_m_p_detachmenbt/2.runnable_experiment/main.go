package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("A running")
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println("B running")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(6 * time.Second)
}
