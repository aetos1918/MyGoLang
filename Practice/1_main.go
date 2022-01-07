package main

import (
	"fmt"
)

const aConst int = 45

func main_1() {

	// Explicit variable assign
	var aString string = "Hello hacker ..!"
	fmt.Println(aString)
	fmt.Printf("The variable type is: %T\n", aString)

	var anInteger int = 45
	fmt.Println(anInteger)
	fmt.Printf("The variable type is: %T\n", anInteger)

	var aDefault int
	fmt.Println(aDefault)

	// Implicit variable assign
	var anotherString = "hacker"
	fmt.Println(anotherString)
	fmt.Printf("The variable type is: %T\n", anotherString)

	d := 445
	fmt.Println(d)
	fmt.Printf("The variable type is: %T\n", d)

	fmt.Println(aConst)
}
