package main

import "fmt"

type Mover interface {
	Move() string
}

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (Dog) Move() string  { return "running" }
func (Dog) Speak() string { return "woof" }

type Truck struct{}

func (Truck) Move() string  { return "driving" }
func (Truck) Speak() string { return "horn" }

func Emit(s Speaker) {
	fmt.Println("Sound:", s.Speak())
}

func Travel(m Mover) {
	fmt.Println("Movement:", m.Move())
}

func main() {
	d := Dog{}
	t := Truck{}

	Emit(d)
	Travel(d)

	Emit(t)
	Travel(t)
}
