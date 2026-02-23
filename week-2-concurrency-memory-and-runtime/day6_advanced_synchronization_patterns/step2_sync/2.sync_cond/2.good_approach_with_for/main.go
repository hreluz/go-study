package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	mu   sync.Mutex
	cond *sync.Cond
	data []int
}

func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *Queue) Push(v int) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()

	// Signal AFTER modifying shared state.
	q.cond.Signal()
}

func (q *Queue) Pop() int {
	q.mu.Lock()
	for len(q.data) == 0 { // MUST be for, not if
		q.cond.Wait() // atomically: unlock mu, sleep, then lock mu again
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}

func main() {
	q := NewQueue()

	// Consumer
	go func() {
		for {
			v := q.Pop()
			fmt.Println("got", v)
		}
	}()

	// Producer
	for i := 0; i < 5; i++ {
		fmt.Println("push", i)
		q.Push(i)
		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(1 * time.Second)
}
