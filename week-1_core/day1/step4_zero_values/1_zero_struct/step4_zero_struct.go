package main

import "fmt"

type Stats struct {
	Count int
	Sum   int
}

func (s *Stats) Add(x int) {
	s.Count++
	s.Sum += x
}

func main() {
	var s Stats // zero value
	s.Add(10)
	s.Add(5)
	fmt.Println(s.Count, s.Sum)
}
