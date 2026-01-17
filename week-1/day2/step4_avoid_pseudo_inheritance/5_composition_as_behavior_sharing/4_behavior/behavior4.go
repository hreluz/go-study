package main

import "fmt"

// Interface defined by behavior
type Speaker interface {
	Speak() string
}

// Type 1: Dog has Speak() because we implement it
type Dog struct{}

func (Dog) Speak() string { return "Woof!" }

// Type 2: Robot has Speak() because we implement it
type Robot struct{}

func (Robot) Speak() string { return "Beep-boop" }

// Function accepts any Speaker
func Announce(s Speaker) {
	fmt.Println("Announcement:", s.Speak())
}

func main() {
	Announce(Dog{})
	Announce(Robot{})
}

/**
	There is no hierarchy.
	Dog and Robot are unrelated.
	They simply have the behavior that Speaker requires.
**/
