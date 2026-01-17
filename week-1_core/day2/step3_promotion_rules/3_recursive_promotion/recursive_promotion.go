package main

import "fmt"

type Sensor struct {
	Temperature float64
}

type Engine struct {
	Sensor
}

type Car struct {
	Engine
}

func main() {
	c := Car{
		Engine: Engine{
			Sensor: Sensor{Temperature: 97.5},
		},
	}

	// Field originates in Sensor but promoted twice:
	// Car → Engine → Sensor
	fmt.Println("Temperature:", c.Temperature)
}
