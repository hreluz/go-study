package main

import "fmt"

type Counter struct {
	N int
}

func incPtr(c *Counter) {
	c.N++
}

func main() {
	c := Counter{N: 10}
	incPtr(&c)
	fmt.Println("after incPtr:", c.N) // expect 11
}
