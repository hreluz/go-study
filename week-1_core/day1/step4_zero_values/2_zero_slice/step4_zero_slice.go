package main

import "fmt"

type Bag struct {
	items []string
}

func (b *Bag) Add(x string) {
	b.items = append(b.items, x)
}

func main() {
	var b Bag // zero value: b.items == nil
	b.Add("a")
	b.Add("b")
	fmt.Println(b.items)
}
