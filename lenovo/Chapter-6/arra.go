package main

import (
	"fmt"
)

func main() {
	s := make([]int, 4)

	fmt.Println(s)

	s2 := []int{1, 2, 3, 4, 5}

	fmt.Println(s2)

	s2 = append(s2, 7, 8, 9) // append values to a slice

	fmt.Println(s2)

	fmt.Println(len(s2))

	fmt.Println(cap(s2))

	nums1 := [5]int{1, 2, 3, 4, 5}

	nums2 := nums1

	nums2[0] = 45

	fmt.Println(nums1) // 1 2 3 4 5
	fmt.Println(nums2) // 45 2 3 4 5

	t := []int{1, 5, 7, 8, 9}

	nums3 := make([]int, len(t))
	copy(nums3, t)

	fmt.Println(nums3)

	fmt.Println(nums3)

}
