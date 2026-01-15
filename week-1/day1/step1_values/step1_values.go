package main

import "fmt"

type Counter struct {
	N int
}

func incValue(c Counter) {
	c.N++
	fmt.Printf("%p\n", &c)
}

func main() {
	c := Counter{N:10}
	incValue(c)
	fmt.Println("after incValue:", c.N)
	fmt.Printf("%p\n", &c)
}
