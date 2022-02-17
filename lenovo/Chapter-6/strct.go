package main

import (
	"fmt"
)

type point struct {
	x float32
	y float32
	z float32
}

func main() {
	pt1 := point{x: 2.3, y: 4.5, z: 3.4}

	fmt.Println(pt1)

	pt2 := newPoint(3.4, 5.66, 3.7)

	fmt.Println(pt2)
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}
