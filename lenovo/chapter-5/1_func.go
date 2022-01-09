package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6

	fmt.Println("before swap", x, y)
	// swap1(x, y)
	swap2(&x, &y)
	fmt.Println("After swap", x, y)
	// num := [5]int{1, 2, 3, 4, 5}
	fmt.Println(addNum(1, 2, 3, 4, 5))
}

func swap1(x, y int) {
	x, y = y, x
}

func swap2(a, b *int) {
	*a, *b = *b, *a
}

func addNum(sum int, x ...int) int {
	// sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
