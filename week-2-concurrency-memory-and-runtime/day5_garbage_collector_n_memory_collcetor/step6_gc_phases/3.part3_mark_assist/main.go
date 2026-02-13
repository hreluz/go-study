// main.go
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func churnAlloc(iters int) {
	for i := 0; i < iters; i++ {
		_ = make([]byte, 1024) // frequent allocations
		if i%50_000 == 0 {
			runtime.Gosched() // yield to increase interleaving
		}
	}
}

func main() {
	// More frequent GC to increase chance allocations overlap with marking
	debug.SetGCPercent(20)

	fmt.Println("Start churn; run with tracing enabled:")
	fmt.Println("  GODEBUG=gctrace=1,gcpacertrace=1 go run main.go")

	// Allocate a lot. This tends to trigger GC cycles and assist work.
	churnAlloc(2_000_000)

	// Force one last GC for a final trace line.
	runtime.GC()
	fmt.Println("Done")
}

//The main function itself runs inside a goroutine.

/**
	When GC enters mark phase:

		The runtime launches background marking workers.

		Meanwhile, your main goroutine continues allocating.

		When it allocates during mark phase, it may be forced to perform mark assist work.

	So assist does not require multiple user-created goroutines.

		It only requires:

			A goroutine allocating

			While GC is in mark phase

	Even a single goroutine program can trigger mark assist.

	Important Mental Model

		There are two types of goroutines here:

			1.Your application goroutines (like main)

			2.Runtime GC worker goroutines

		Assist means:
			Your application goroutine temporarily helps the GC workers.


GC enters MARK phase
        |
        v
main goroutine calls make(...)
        |
        v
runtime says: before you allocate, help mark some objects
**/
