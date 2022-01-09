package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}

	max := 100
	a, b := 0, 1
	fmt.Printf("\n%v ", a)
	for b < max {
		fmt.Printf("%v ", b)
		a, b = b, a+b
	}

	// infinite()
	iterate()
}

func infinite() {
	for {
		fmt.Println("Enter QUIT to exit")
		var input string
		fmt.Print("Please enter a string:")
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "QUIT" {
			break
		}
	}
}

func iterate() {
	names := [4]string{"Hacker", "Killer", "Psycho"} // array

	for i, v := range names {
		fmt.Printf("%v\t%v\n", i, v)
	}

	for pos, char := range "Hello World" {
		fmt.Println(pos, char)
	}

	for pos, char := range "Hello World" {
		fmt.Printf("%d %c\n", pos, char)
	}
}
