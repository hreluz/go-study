package main

import "fmt"

// Interfaces express required BEHAVIOR, not hierarchy.
type Speaker interface {
	Speak() string
}

type Mover interface {
	Move() string
}

// Data-only component
type Identity struct {
	Name string
}

// Behavior components
type SoundMaker struct {
	Sound string
}

func (s SoundMaker) Speak() string {
	return s.Sound
}

type Walker struct {
	SpeedKmH int
}

func (w Walker) Move() string {
	return fmt.Sprintf("walking at %d km/h", w.SpeedKmH)
}

type Wheels struct {
	SpeedKmH int
}

func (w Wheels) Move() string {
	return fmt.Sprintf("rolling at %d km/h", w.SpeedKmH)
}

// Concrete types COMPOSE data + behaviors

type Dog struct {
	Identity
	SoundMaker
	Walker
}

type Cat struct {
	Identity
	SoundMaker
	Walker
}

type Robot struct {
	Identity
	SoundMaker
	Wheels
}

// Functions operate on behavior, not concrete types

func DescribeSpeaker(s Speaker) {
	fmt.Println("Speaks:", s.Speak())
}

func DescribeMover(m Mover) {
	fmt.Println("Moves:", m.Move())
}

func main() {
	dog := Dog{
		Identity:   Identity{Name: "Fido"},
		SoundMaker: SoundMaker{Sound: "Woof!"},
		Walker:     Walker{SpeedKmH: 10},
	}

	cat := Cat{
		Identity:   Identity{Name: "Misu"},
		SoundMaker: SoundMaker{Sound: "Meow"},
		Walker:     Walker{SpeedKmH: 5},
	}

	robot := Robot{
		Identity:   Identity{Name: "R2"},
		SoundMaker: SoundMaker{Sound: "Beep-boop"},
		Wheels:     Wheels{SpeedKmH: 20},
	}

	fmt.Println("Dog:", dog.Name)
	DescribeSpeaker(dog)
	DescribeMover(dog)

	fmt.Println("\nCat:", cat.Name)
	DescribeSpeaker(cat)
	DescribeMover(cat)

	fmt.Println("\nRobot:", robot.Name)
	DescribeSpeaker(robot)
	DescribeMover(robot)
}
