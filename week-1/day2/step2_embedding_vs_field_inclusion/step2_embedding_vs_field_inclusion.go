package main

type Engine struct {
	Horsepower int
}

func (e Engine) Power() int {
	return e.Horsepower
}

// Version A: normal field inclusion
type CarA struct {
	Brand  string
	Engine Engine
}

// Version B: embedding
type CarB struct {
	Brand string
	Engine
}

func main() {
	var a CarA
	a.Engine.Horsepower = 150
	_ = a.Engine.Power()

	var b CarB
	b.Horsepower = 150
	_ = b.Power()
}
