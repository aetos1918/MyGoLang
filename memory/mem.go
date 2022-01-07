package main

import (
	"fmt"
	"sort"
)

func main() {
	// hack := new(map[string]int)
	// hack["sree"] = 25

	// maps()
	// mem()
	// arr()
	// slic()
	// gomap()
	strt()
}

func maps() {
	hack := make(map[string]int)
	hack["sree"] = 25
	fmt.Println("My map: ", hack)
}

func mem() {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p: ", *p)

	value1 := 53
	ptr1 := &value1
	fmt.Println("Ptr1 value: ", *ptr1)

	*ptr1 = *ptr1 / 12

	fmt.Println("Ptr1 value: ", *ptr1)
	fmt.Println("Ptr1 value: ", value1)
}

func arr() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("Arrays: ", colors)

	var numbers = [5]int{5, 6, 4, 3, 2}
	fmt.Println(numbers)

	fmt.Println("Length of colors: ", len(colors))
	fmt.Println("Length of numbers: ", len(numbers))
}

func slic() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println("colors: ", colors)

	colors = append(colors, "Purple")
	fmt.Println("Arrays: ", colors)

	colors = append(colors[1:])
	fmt.Println("Arrays: ", colors)

	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 45
	numbers[2] = 56
	numbers[3] = 67
	numbers[4] = 98
	// numbers[5] = 98

	fmt.Println("Numbers: ", numbers)
	numbers = append(numbers, 54)
	fmt.Println("Numbers: ", numbers)
	sort.Ints(numbers)
	fmt.Println("Numbers: ", numbers)
}

func gomap() {
	states := make(map[string]string)

	states["AP"] = "Andhra Pradesh"
	states["KA"] = "Karnataka"
	states["TN"] = "Tamilnadu"
	states["TS"] = "Telangana"
	states["OD"] = "Odissa"
	states["WB"] = "West Bengal"
	fmt.Println(states)
	delete(states, "KA")
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	fmt.Println(keys)

	for j := range keys {
		fmt.Printf("%v: %v\n", keys[j], states[keys[j]])
	}
}
