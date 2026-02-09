package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		x = 3
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(x)
}
