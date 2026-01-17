package main

import "fmt"

func appendValue(s []int) {
	s = append(s, 100)
	fmt.Println("inside:", s)
}

func main() {
	a := []int{1, 2, 3}
	appendValue(a)
	fmt.Println("outside:", a)
}
