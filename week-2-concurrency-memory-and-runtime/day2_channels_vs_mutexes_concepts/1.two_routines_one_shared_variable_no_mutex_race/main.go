package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0

	go func() {
		for i := 0; i < 100000; i++ {
			count++
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			count++
		}
	}()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("count:", count)
}

