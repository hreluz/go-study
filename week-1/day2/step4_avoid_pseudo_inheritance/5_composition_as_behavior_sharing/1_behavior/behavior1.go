package main

import "fmt"

// Behavior component
type SoundMaker struct {
	Sound string
}

func (s SoundMaker) Speak() string {
	return s.Sound
}

// Types that *have* this behavior
type Dog struct {
	SoundMaker
}

type Robot struct {
	SoundMaker
}

func main() {
	dog := Dog{SoundMaker{Sound: "Woof!"}}
	robot := Robot{SoundMaker{Sound: "Beep-boop"}}

	fmt.Println("Dog:", dog.Speak())
	fmt.Println("Robot:", robot.Speak())
}
