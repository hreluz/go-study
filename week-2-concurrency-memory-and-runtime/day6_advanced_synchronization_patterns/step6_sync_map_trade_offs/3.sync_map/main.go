package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const (
		keys       = 1000
		goroutines = 200
		opsPerG    = 20000
		writeEvery = 100 // 1% writes
	)

	var sm sync.Map

	// Initialize
	for i := 0; i < keys; i++ {
		sm.Store(i, i)
	}

	start := time.Now()
	var wg sync.WaitGroup

	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(seed int64) {
			defer wg.Done()
			r := rand.New(rand.NewSource(seed))
			for i := 0; i < opsPerG; i++ {
				k := r.Intn(keys)
				if i%writeEvery == 0 {
					sm.Store(k, r.Int())
				} else {
					_, _ = sm.Load(k)
				}
			}
		}(int64(g + 1))
	}

	wg.Wait()
	fmt.Println("sync.Map:", time.Since(start))
}
