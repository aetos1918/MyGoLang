package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello From Go"
	file, err := os.Create("./hacker.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a line with %v characters.", length)
	defer file.Close()
	defer readFile("./hacker.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("\nText read from file: ", string(data))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
