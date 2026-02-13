package main

import (
	"fmt"
	"runtime"
)

func good() {
	buf := make([]byte, 1024)
	for i := 0; i < 1_000_000; i++ {
		_ = buf
	}
}

func main() {
	good()
	runtime.GC()
	fmt.Println("done")
}

/**
	Only one allocation.

	Massively less GC pressure.
**/
