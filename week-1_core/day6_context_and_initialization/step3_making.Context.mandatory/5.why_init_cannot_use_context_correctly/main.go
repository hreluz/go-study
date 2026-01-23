package main

import (
	"context"
	"fmt"
	"time"
)

func init() {
	go func() {
		// No caller, no deadline, no cancellation
		time.Sleep(1 * time.Second)
		fmt.Println("init work finished (uncontrolled)")
	}()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	<-ctx.Done()
	fmt.Println("Main context canceled")
	time.Sleep(2 * time.Second)
}
