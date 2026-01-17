package main

import "fmt"

type Speaker interface {
	Speak() string
}

type SoundMaker struct {
	Sound string
}

func (s SoundMaker) Speak() string {
	return s.Sound
}

type Dog struct {
	SoundMaker
}

type CarAlarm struct {
	SoundMaker
}

func main() {
	dog := Dog{SoundMaker{"Woof!"}}
	alarm := CarAlarm{SoundMaker{"BEEP BEEP"}}

	var s Speaker

	s = dog
	fmt.Println("Dog says:", s.Speak())

	s = alarm
	fmt.Println("Alarm says:", s.Speak())
}

/**
	Still no hierarchies.
	SoundMaker provides behavior that multiple types reuse.
**/
