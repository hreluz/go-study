package main

import "fmt"

type Identity struct {
	Name string
}

type SoundMaker struct {
	Sound string
}

func (s SoundMaker) Speak() string {
	return s.Sound
}

type Dog struct {
	Identity
	SoundMaker
}

func main() {
	d := Dog{
		Identity:   Identity{Name: "Bobby"},
		SoundMaker: SoundMaker{Sound: "Woof!"},
	}

	fmt.Println("Name:", d.Name)
	fmt.Println("Speak:", d.Speak())
}
