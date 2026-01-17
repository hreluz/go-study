package main

import "fmt"

// Pure data: identity only
type Identity struct {
	Name string
}

// Behavior component: sound
type SoundMaker struct {
	Sound string
}

func (s SoundMaker) Speak() string {
	return s.Sound
}

// Behavior component: movement (example of additional behavior)
type Walker struct {
	SpeedKmH int
}

func (w Walker) Move() string {
	return fmt.Sprintf("walking at %d km/h", w.SpeedKmH)
}

// Dog COMPOSES behaviors; it is not "inheriting"
type Dog struct {
	Identity
	SoundMaker
	Walker
}

// Cat COMPOSES behaviors as well
type Cat struct {
	Identity
	SoundMaker
	Walker
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

	fmt.Println("Dog:", dog.Name, "-", dog.Speak(), "-", dog.Move())
	fmt.Println("Cat:", cat.Name, "-", cat.Speak(), "-", cat.Move())
}
