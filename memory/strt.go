package main

import (
	"fmt"
)

func strt() {
	poodle := Dog{"Poodle", 55}
	fmt.Println("Structs: ", poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("%v\t%v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 42
	fmt.Printf("%v\t%v\n", poodle.Breed, poodle.Weight)
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}
