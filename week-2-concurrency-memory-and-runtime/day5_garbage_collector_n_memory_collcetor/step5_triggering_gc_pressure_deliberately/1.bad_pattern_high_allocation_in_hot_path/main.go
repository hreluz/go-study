package main

import (
	"fmt"
	"runtime"
)

func bad() {
	for i := 0; i < 1_000_000; i++ {
		_ = make([]byte, 1024)
	}
}

func main() {
	bad()
	runtime.GC()
	fmt.Println("done")
}

/**
	This creates massive allocation churn.

		Effects:

		High allocation rate

		More GC cycles

		More mark assist

		Increased CPU overhead
**/
