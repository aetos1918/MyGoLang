package main

import (
	"fmt"
)

func main() {
	num := [5]int{1, 2, 3, 4, 5}
	gf := filter(num)

	fmt.Println(gf(num))
}

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for _, val := range err {
		if cond(val) {
			result = append(result, val)
		}
	}
	return result
}
