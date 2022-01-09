package main

import (
	"fmt"
)

func main() {
	// nestloop()
	// breakinner()
	breakouter()
}

func nestloop() {
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Printf("%v * %v = %v\n", i, j, i*j)
		}
		fmt.Println("---------------")
	}
}

func breakinner() {
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if i == 3 {
				break
			}
			fmt.Printf("%v * %v = %v\n", i, j, i*j)
		}
		fmt.Println("---------------")
	}
}

func breakouter() {
Outerloop:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break Outerloop
				// continue Outerloop
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("-----------")
	}
}
