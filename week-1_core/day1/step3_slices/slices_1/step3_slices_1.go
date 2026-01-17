package main

import "fmt"

func changeFirst(s []int) {
	s[0] = 99
}

func main() {
	a := []int{1, 2, 3}
	changeFirst(a)
	fmt.Println(a)
}
