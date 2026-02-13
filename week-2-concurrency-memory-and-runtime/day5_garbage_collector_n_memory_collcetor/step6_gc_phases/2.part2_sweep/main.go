// part2_sweep.go
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
	debug.SetGCPercent(50)

	printMem("start")

	live := makeList(200_000)
	runtime.KeepAlive(live)

	printMem("after build list")

	fmt.Println("forcing GC with reachable graph...")
	runtime.GC()
	runtime.KeepAlive(live)

	printMem("after GC (reachable)")

	// 🔴 DROP THE ROOT
	live = nil

	fmt.Println("forcing GC after dropping reference...")
	runtime.GC()

	printMem("after GC (unreachable)")
}

/**
	What should happen

		Before live = nil:

			The graph is reachable → GC marks it → memory remains high.

		After live = nil:

			No root points to the list.

			The graph becomes unreachable.

			Next GC marks nothing in that chain.

			Sweep reclaims memory.

			HeapAlloc and HeapObjects should drop significantly.

	After live = nil, the linked list is no longer reachable from any GC root, so it becomes eligible for sweep.
**/
