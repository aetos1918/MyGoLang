package main

import (
	"cmp"
	"fmt"
)

type Point struct {
	X    float32
	Y    float32
	Z    float32
	Name []string
}

func main() {
	pt1 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt1"}}
	pt2 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt"}}
	pt3 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt"}}
	fmt.Println(cmp.Equal(pt1, pt2)) // false
	fmt.Println(cmp.Equal(pt2, pt3)) // true
}
