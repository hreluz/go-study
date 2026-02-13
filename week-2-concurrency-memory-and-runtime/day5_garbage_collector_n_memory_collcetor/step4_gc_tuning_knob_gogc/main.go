package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func stats(label string) runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: HeapAlloc=%d KB, HeapObjects=%d, NumGC=%d\n",
		label, m.HeapAlloc/1024, m.HeapObjects, m.NumGC)
	return m
}

func main() {
	debug.SetGCPercent(400) // try 20, 100, 200

	stats("start")

	var peak uint64
	var junk [][]byte

	for i := 0; i < 400_000; i++ {
		junk = append(junk, make([]byte, 1024)) // keep everything reachable

		// sample occasionally
		if i%20_000 == 0 {
			m := stats(fmt.Sprintf("i=%d", i))
			if m.HeapAlloc > peak {
				peak = m.HeapAlloc
			}
		}
	}

	// Do NOT force runtime.GC() here; we want to observe the natural peak behavior.
	fmt.Printf("PEAK HeapAlloc=%d KB\n", peak/1024)

	_ = junk // keep reachable
}
