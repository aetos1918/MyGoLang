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

	var pt1 point

	pt1.x = 4.3
	pt1.y = 3.4
	pt1.z = 5.6

	fmt.Println(pt1)

	pt2 := point{x: 5.6, y: 4.7, z: 34.78}

	fmt.Println(pt2)

	pt4 := newPoint(7.8, 3.2, 4.5)

	fmt.Println(pt4)

	pt5 := pt4

	pt5.x = 0 // changes will happen in both

	fmt.Println(pt4)
	fmt.Println(pt5)

	pt6 := *pt4
	pt6.x = 25.3 // changes will happen in one
	fmt.Println(pt6)
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}
