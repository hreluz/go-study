package main

import "fmt"

func goroutineEscape() {
	x := 10
	go func() {
		fmt.Println(x)
	}()
}

func syncNoEscape() {
	x := 10
	func() {
		fmt.Println(x)
	}()
}

func main() {
	goroutineEscape()
	syncNoEscape()
}

// go run -gcflags="-m -m" main.go
