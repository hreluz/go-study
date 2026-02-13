package main

type Data struct {
	x int
}

func noEscape() {
	d := Data{x: 42}
	_ = d.x
}

func escape() *Data {
	d := Data{x: 42}
	return &d
}

func main() {
	noEscape()
	_ = escape()
}

// go run -gcflags="-m -m" main.go
