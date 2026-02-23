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

// Pop with WRONG pattern: uses if
func (q *Queue) PopBroken(id string) int {
	q.mu.Lock()
	if len(q.data) == 0 {
		fmt.Println(id, "waiting...")
		q.cond.Wait()
	}
	// ❗ Danger: assumes data exists
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	fmt.Println(id, "got", v)
	return v
}

func (q *Queue) Push(v int) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.cond.Signal()
	q.mu.Unlock()
}

func main() {
	q := NewQueue()

	// Two consumers
	go func() {
		time.Sleep(100 * time.Millisecond)
		q.PopBroken("Consumer A")
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		q.PopBroken("Consumer B")
	}()

	// Producer pushes only ONE item
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Producer pushing 1 item")
	q.Push(42)

	time.Sleep(1 * time.Second)
}
