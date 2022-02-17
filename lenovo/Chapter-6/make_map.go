package main

import (
	"fmt"
)

func main() {
	heights := make(map[string]int)

	heights["Suresh"] = 112

	fmt.Println(heights)

	fmt.Println(heights["Naresh"]) // 0

	area := map[string]int{
		"Hyd": 1200,
		"Sec": 700,
		"Vzg": 600, // <-- Note the comma here
	}

	fmt.Println(area)

	key := "Hyd"

	if area[key] != 0 {
		fmt.Printf("%v presnt in area", key)
	}
}
