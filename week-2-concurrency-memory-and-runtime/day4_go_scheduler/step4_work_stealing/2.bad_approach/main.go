package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	var done int32
	var count0 int64
	var count1 int64

	// Producer that tends to enqueue work on one P.
	go func() {
		runtime.LockOSThread() // tends to keep this goroutine on one M/P pairing
		for atomic.LoadInt32(&done) == 0 {
			go func() {
				// tiny work
				atomic.AddInt64(&count0, 1)
			}()
		}
	}()

	// Another goroutine doing CPU-bound work (simulating other load).
	go func() {
		for atomic.LoadInt32(&done) == 0 {
			atomic.AddInt64(&count1, 1)
		}
	}()

	time.Sleep(2 * time.Second)
	atomic.StoreInt32(&done, 1)
	time.Sleep(200 * time.Millisecond)

	fmt.Println("count0 (spawned tiny tasks):", atomic.LoadInt64(&count0))
	fmt.Println("count1 (busy loop):       ", atomic.LoadInt64(&count1))
}

/**
	Why this is “bad” conceptually

	The design creates a storm of tiny goroutines from one place.

	This increases scheduling overhead and can create temporary imbalance.

	Work stealing may mitigate it, but you are fighting the scheduler.

	Real-life analogy:

	One counter keeps generating new customers for only its own line, forcing the other counter to repeatedly “steal” customers instead of having customers distribute naturally.
**/
