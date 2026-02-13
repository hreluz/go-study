package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 100_000; i++ {
		_ = make([]byte, 1024)
	}

	runtime.GC()

	fmt.Println("Allocated but not kept")
}
