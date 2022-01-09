package main

import (
	"fmt"
)

func main() {
	var num int = 5
	// fmt.Printf("Enter num: ")
	// fmt.Scanf("%d", &num)
	var dayOfWeek string
	switch num {
	case 1:
		dayOfWeek = "Sunday"
	case 2:
		dayOfWeek = "Monday"
	case 3:
		dayOfWeek = "Tuesday"
	case 4:
		dayOfWeek = "Wednesday"
	case 5:
		dayOfWeek = "Thursday"
	case 6:
		dayOfWeek = "Friday"
	case 7:
		dayOfWeek = "Saturday"
	default:
		dayOfWeek = "__Error__"
	}
	fmt.Println(dayOfWeek)

	grade := "C"
	switch grade {
	case "A":
		fallthrough
	case "B":
		fallthrough
	case "C":
		fallthrough
	case "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Absent")
	}

	switch grade {
	case "A", "B", "C", "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Absent")
	}

	/*
		Below construct makes it easy to perform conditional checks within the case
		expressions.
	*/
	score := 79
	grade = ""
	switch {
	case score < 50:
		grade = "F"
	case score < 60:
		grade = "D"
	case score < 70:
		grade = "C"
	case score < 80:
		grade = "B"
	default:
		grade = "A"
	}
	fmt.Println(grade) // B

}
