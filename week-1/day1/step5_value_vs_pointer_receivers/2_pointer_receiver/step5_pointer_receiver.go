package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{1, 2}
	p.Move(10, 10)
	fmt.Println(p)
}
