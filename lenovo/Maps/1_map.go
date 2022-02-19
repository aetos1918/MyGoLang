package main

import (
	"fmt"
)

func main() {
	var height = make(map[string]int)
	height["sree"] = 5

	fmt.Println(height)

	news := map[string]int{
		"Peter": 170,
		"Joan":  168,
		"Jan":   175, // <-- note the comma here
	}

	for k, v := range news {
		fmt.Printf("%v\t%v\n", k, v)
	}

	// Checking the existence of key
	if v, ok := news["jim"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist..!")
	}

	// Deleting a Key
	if _, ok := news["hack"]; ok {
		delete(news, "hack")
	} else {
		fmt.Println("Key does not exist to delete")
	}

}
