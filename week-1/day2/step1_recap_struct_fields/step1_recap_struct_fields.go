package main

type Engine struct {
	Horsepower int
}

type Car struct {
	Brand  string
	Engine Engine
}
