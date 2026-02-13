// part1_mark.go
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type Node struct {
	next *Node
	pad  [256]byte
}

func makeList(n int) *Node {
	var head *Node
	for i := 0; i < n; i++ {
		head = &Node{next: head}
	}
	return head
}

func printMem(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%-24s HeapAlloc=%7d KB  HeapObjects=%7d  NumGC=%d\n",
		label, m.HeapAlloc/1024, m.HeapObjects, m.NumGC)
}

func main() {
	// Make GC run more often (easier to observe in short programs)
	debug.SetGCPercent(50)

	printMem("start")

	// Create a large reachable object graph
	live := makeList(200_000)

	// Keep it live up to this point (avoid compiler liveness optimizations)
	runtime.KeepAlive(live)

	printMem("after build list")

	// Force a GC while 'live' is reachable => GC must MARK it
	fmt.Println("forcing GC...")
	runtime.GC()

	// Keep it alive through the GC and print (extra safety)
	runtime.KeepAlive(live)

	printMem("after GC (reachable)")
}

// GODEBUG=gctrace=1 go run part1_mark.go
/**
	live is a stack variable in main.

	At the moment runtime.GC() runs, live is live.

	Therefore GC scans it as a root.

	From live, it follows next pointers.

	The entire linked list becomes reachable.

	Therefore it must be marked and cannot be swept.
**/
