package main

import (
	"context"
	"fmt"
	"time"
)

type Service struct {
	ctx context.Context
}

func NewService(ctx context.Context) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) DoWork() {
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Work done")
	case <-s.ctx.Done():
		fmt.Println("Canceled:", s.ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	service := NewService(ctx)

	cancel() // First request ends

	// Second request wrongly reuses canceled context
	service.DoWork()
}
