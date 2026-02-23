package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RWMap struct {
	mu sync.RWMutex
	m  map[int]int
}

func NewRWMap() *RWMap {
	return &RWMap{m: make(map[int]int)}
}

func (r *RWMap) Load(k int) (int, bool) {
	r.mu.RLock()
	v, ok := r.m[k]
	r.mu.RUnlock()
	return v, ok
}

func (r *RWMap) Store(k, v int) {
	r.mu.Lock()
	r.m[k] = v
	r.mu.Unlock()
}

func main() {
	const (
		keys       = 1000
		goroutines = 200
		opsPerG    = 20000
		writeEvery = 100 // 1% writes
	)

	// Initialize
	rwm := NewRWMap()
	for i := 0; i < keys; i++ {
		rwm.Store(i, i)
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
					rwm.Store(k, r.Int())
				} else {
					_, _ = rwm.Load(k)
				}
			}
		}(int64(g + 1))
	}

	wg.Wait()
	fmt.Println("map+RWMutex:", time.Since(start))
}
