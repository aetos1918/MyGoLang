// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

func main() {
	slc := []int{1, 2, 3, 45, 6}
	fmt.Println(slc)
	slc, err := delete(slc, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(slc)
	}

}

func delete(arr []int, index int) ([]int, error) {
	if index < 0 || index >= len(arr) {
		return nil, errors.New("New Error came")
	}
	arr = append(arr[:index], arr[index+1:]...)
	return arr, nil
}
