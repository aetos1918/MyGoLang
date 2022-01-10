package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 23, "Woof!"}
	poodle.Speak()
	poodle.Sound = "Arf..!"
	poodle.SpeakThreeTimes()
}

// Dog is a struct
type Dog struct {
	Breed  string
	weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
