package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Value 1: ")
	input_1, _ := reader.ReadString('\n')
	value_1, err_1 := strconv.ParseFloat(strings.TrimSpace(input_1), 64)

	if err_1 != nil {
		panic("this is code Error")
	}

	fmt.Printf("Value 2: ")
	input_2, _ := reader.ReadString('\n')
	value_2, err_2 := strconv.ParseFloat(strings.TrimSpace(input_2), 64)

	if err_2 != nil {
		panic("this is code Error")
	}

	final := value_1 + value_2
	final = math.Round(final*100) / 100

	fmt.Printf("sum of %v and %v is: %v", value_1, value_2, final)

}
