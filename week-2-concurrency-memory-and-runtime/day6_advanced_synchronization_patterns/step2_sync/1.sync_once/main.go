package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func expensiveSetup() {
	fmt.Println("setup start")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("setup done")
}

func Init() {
	once.Do(expensiveSetup)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Init()
			fmt.Println("goroutine", id, "passed Init")
		}(i)
	}
	wg.Wait()
}
