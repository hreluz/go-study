package main

import (
	"fmt"
	"time"
)

func worker() {
	fmt.Println("worker: start")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("worker: end (return)")
}

func main() {
	fmt.Println("main: before go")
	go worker()

	fmt.Println("main: after go")
	time.Sleep(1 * time.Second) // keep process alive long enough to see worker end
	fmt.Println("main: exit")
}

