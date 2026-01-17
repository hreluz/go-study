package main

import "fmt"

type SoundMaker struct {
	Sound string
}

func (s SoundMaker) Speak() string {
	return s.Sound
}

type Walker struct {
	Steps int
}

func (w Walker) Walk() string {
	return fmt.Sprintf("Walking %d steps", w.Steps)
}

type Dog struct {
	SoundMaker
	Walker
}

type Cat struct {
	SoundMaker
	Walker
}

func main() {
	dog := Dog{
		SoundMaker: SoundMaker{Sound: "Woof!"},
		Walker:     Walker{Steps: 10},
	}

	cat := Cat{
		SoundMaker: SoundMaker{Sound: "Meow"},
		Walker:     Walker{Steps: 5},
	}

	fmt.Println("Dog:", dog.Speak(), "-", dog.Walk())
	fmt.Println("Cat:", cat.Speak(), "-", cat.Walk())
}
