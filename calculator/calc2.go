package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Printf("Enter Value_1: ")
	value_1 := getInput()
	fmt.Printf("Enter Value_2: ")
	value_2 := getInput()
	fmt.Println("Select the operation (+ * - /) :")
	test, _ := reader.ReadString('\n')
	test = strings.TrimSpace(test)
	var result float64
	switch test {
	case "+":
		result = addition(value_1, value_2)
	case "-":
		result = subtraction(value_1, value_2)
	case "*":
		result = multiply(value_1, value_2)
	case "/":
		result = division(value_1, value_2)
	default:
		fmt.Printf(test)
		panic("Invalid Operation ")
	}
	fmt.Printf("the %v operation b/n %v and %v is %v", test, value_1, value_2, result)
}

func getInput() float64 {
	inp, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(inp), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", inp)
		panic(message)
	}
	return value
}
func addition(val1, val2 float64) float64 {
	return val1 + val2
}

func subtraction(val1, val2 float64) float64 {
	return val1 - val2
}

func multiply(val1, val2 float64) float64 {
	return val1 * val2
}

func division(val1, val2 float64) float64 {
	return val1 / val2
}
