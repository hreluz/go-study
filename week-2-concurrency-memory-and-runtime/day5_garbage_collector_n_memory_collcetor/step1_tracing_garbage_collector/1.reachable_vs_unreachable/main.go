package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type Node struct {
	next *Node
	data [1024]byte
}

func printMem(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: HeapAlloc=%d KB, HeapObjects=%d, HeapInuse=%d KB, HeapIdle=%d KB, NumGC=%d\n",
		label,
		m.HeapAlloc/1024,
		m.HeapObjects,
		m.HeapInuse/1024,
		m.HeapIdle/1024,
		m.NumGC,
	)
}

func makeChain(n int) *Node {
	var head *Node
	for i := 0; i < n; i++ {
		head = &Node{next: head}
	}
	return head
}

func touch(head *Node) int {
	sum := 0
	for p := head; p != nil; p = p.next {
		sum += int(p.data[0])
	}
	return sum
}

func main() {
	debug.SetGCPercent(100)
	printMem("start")

	kept := makeChain(50_000)
	fmt.Println("touch sum:", touch(kept))

	// Force a GC while 'kept' is guaranteed live.
	runtime.GC()
	printMem("after GC with kept reachable")
	runtime.KeepAlive(kept) // guarantees liveness through the GC + printMem

	// Now drop the root and GC again.
	kept = nil
	runtime.GC()
	printMem("after GC with kept=nil (unreachable)")
}

/**
go run main.go
	start: HeapAlloc=65 KB, HeapObjects=159, HeapInuse=312 KB, HeapIdle=3464 KB, NumGC=0
	touch sum: 0
	after GC with kept reachable: HeapAlloc=56316 KB, HeapObjects=50162, HeapInuse=57464 KB, HeapIdle=3624 KB, NumGC=5
	after GC with kept=nil (unreachable): HeapAlloc=67 KB, HeapObjects=163, HeapInuse=320 KB, HeapIdle=60768 KB, NumGC=6
**/
