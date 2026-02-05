package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	var mu sync.Mutex

	go func() {
		for i := 0; i < 100000; i++ {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}()

	time.Sleep(200 * time.Millisecond)
	mu.Lock()
	fmt.Println("count:", count)
	mu.Unlock()
}
