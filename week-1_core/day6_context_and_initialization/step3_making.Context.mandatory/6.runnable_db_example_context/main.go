package main

import (
	"context"
	"fmt"
	"time"
)

type App struct {
	db *DB
}

type DB struct{}

func connect(ctx context.Context) (*DB, error) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("DB connected")
		return &DB{}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func NewApp(ctx context.Context) (*App, error) {
	db, err := connect(ctx)
	if err != nil {
		return nil, err
	}
	return &App{db: db}, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	app, err := NewApp(ctx)
	fmt.Println("App:", app, "Error:", err)
}
