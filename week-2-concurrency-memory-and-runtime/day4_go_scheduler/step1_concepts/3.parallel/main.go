package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	go func() {
		for {
		}
	}()

	go func() {
		for {
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

/**


What this proves:

    With GOMAXPROCS(2), two goroutines can execute at the same time.

    CPU usage will reach ~200% on a dual-core machine.
**/
