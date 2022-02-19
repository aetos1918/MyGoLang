package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	slc := make([]int, 5)
	fmt.Println(slc)

	op, _ := insert(slc, 2, 5)

	fmt.Println(op)
}

func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("Index cannot be less than 0")
	}
	if index >= len(orig) {
		return append(orig, value), nil
	}
	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}
