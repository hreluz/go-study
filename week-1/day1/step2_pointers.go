package main

import "fmt"

type Counter struct {
	N int
}

func incPtr(c* Counter) {
	c.N++

	fmt.Printf("incPtr: c (points to Counter) = %p\n", c)
	fmt.Printf("incPtr: &c (address of local param) = %p\n", &c)
}

func main() {
	c := Counter{N:10}
	fmt.Printf("main: &c (address of Counter) = %p\n", &c)
	incPtr(&c)
	fmt.Println("after incPtr:", c.N)
	fmt.Printf("%p\n", &c)
}
