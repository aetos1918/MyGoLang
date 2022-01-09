package main

import (
	"fmt"
	"reflect"
)

const Killer = "Hack3r"

var unused = "No Error"

func main() {
	var name = "Abc"
	fmt.Println(name)
	var (
		firstName string = "abc"
		age       int    = 25
		lastName  string = "xyz"
	)
	fmt.Println(firstName, age, lastName)
	fmt.Println(Killer)

	quotation := `"Anyone who has never made 
	a mistake has never tried
	any thing new."
	-- Albert Einstein`

	fmt.Println(quotation)
	fmt.Println(reflect.TypeOf(age))

}
