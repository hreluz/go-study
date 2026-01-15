package main

type A struct {
	X int
}

type B struct {
	X int
}

type C struct {
	A
	B
}

func main() {
	var c C

	// ERROR: ambiguous selector c.X
	_ = c.X
}
