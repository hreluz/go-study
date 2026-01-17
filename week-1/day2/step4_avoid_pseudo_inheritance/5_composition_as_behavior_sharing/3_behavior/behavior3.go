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

type Robot struct {
	SoundMaker
}

func announce(s Speaker) {
	fmt.Println("Announcement:", s.Speak())
}

func main() {
	d := Dog{SoundMaker{"Woof!"}}
	r := Robot{SoundMaker{"Beep!"}}

	announce(d)
	announce(r)
}

// A type satisfies an interface by having certain behavior, 	not by being part of a hierarchy.
