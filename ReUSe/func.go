package main

import (
	"fmt"
)

func main() {
	sum := addValues(5, 8)
	fmt.Println("Sum is: ", sum)
	tot, num := addAllValues(1, 2, 3, 4, 6, 2, 1)
	fmt.Println("Total ", tot)
	fmt.Println("LengthL ", num)
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, i := range values {
		total += i
	}
	return total, len(values)
}
