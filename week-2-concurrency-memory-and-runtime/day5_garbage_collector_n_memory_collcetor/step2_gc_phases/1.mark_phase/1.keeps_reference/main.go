package main

import (
	"fmt"
	"runtime"
)

func main() {
	var data [][]byte

	for i := 0; i < 100_000; i++ {
		data = append(data, make([]byte, 1024))
	}

	runtime.GC()

	fmt.Println("Allocated 100k 1KB objects")
}
