package main

import (
	"fmt"
)

func main() {
	i := func() int {
		return 5
	}
	fmt.Println(i())

	gen := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}

}

func fib() func() int { // The fib() function returns a function that returns an int
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)
		return f1
	}
}
