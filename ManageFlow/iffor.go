package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ifst()
	// swtch()
	loop()
}

func ifst() {
	theAns := 100
	var result string

	if theAns < 0 {
		result = "Less Than Zero"
	} else if theAns > 0 && theAns < 50 {
		result = "Less than 50"
	} else {
		result = "Greater than 50"
	}
	fmt.Println(result)

	if x := -42; x < 0 {
		result = "Less Than Zero"
	} else if x == 0 {
		result = "Equal to Zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)
}

func swtch() {
	fmt.Println("Switches")
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day: ", dow)
	var result string
	switch dow {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	default:
		result = "It's Another Day"
	}
	fmt.Println(result)
}

func loop() {
	colors := []string{"Red", "Green", "Blue"}
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, v := range colors {
		fmt.Println(v)
	}

	value := 1

	for value < 10 {
		fmt.Println("Value: ", value)
		value++
	}

	sum := 1

	for sum < 1000 {
		fmt.Println("Sum: ", sum)
		sum += sum
		if sum > 200 {
			// break
			// continue
			goto theEnd
		}
	}
theEnd:
	fmt.Println("End of Program")

}
