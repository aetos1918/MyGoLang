package main

import (
	"fmt"
	"math"
)

func maths() {
	fmt.Println("Maths operations")

	aInt := 34
	aFloat := 3.4567
	// final := aInt + aFloat Invalid Operation
	final := float64(aInt) + aFloat

	fmt.Println("Final number is: ", final)

	rnd := math.Round(final)
	fmt.Println("Round value: ", rnd)

	crnd := math.Round(final*100) / 100
	fmt.Println("Correct Round value: ", crnd)

	circleRadius := 15.5
	circumferance := circleRadius * 2 * math.Pi
	fmt.Printf("Circumferance is: %.2f\n", circumferance)

}
