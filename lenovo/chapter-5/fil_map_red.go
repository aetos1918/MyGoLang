package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5}
	evens := filter(num,
		func(val int) bool {
			return val%2 == 0
		})

	fmt.Println(evens)
}

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for _, val := range arr {
		if cond(val) {
			result = append(result, val)
		}
	}
	return result
}
