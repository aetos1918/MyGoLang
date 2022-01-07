package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func uinp() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered:", input)

	fmt.Println("Enter a number: ")
	numInput, _ := reader.ReadString('\n')

	float, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		fmt.Println("Error is: ", err)
	} else {
		fmt.Println("You entered: ", float)
	}
}
