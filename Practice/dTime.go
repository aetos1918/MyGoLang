package main

import (
	"fmt"
	"time"
)

func dtime() {
	fmt.Println("Dates and Times")
	n := time.Now()

	fmt.Println("time is: ", n)

	t := time.Date(2022, time.January, 22, 3, 0, 12, 0, time.UTC)

	fmt.Println("Date is: ", t)
	fmt.Println(t.Format(time.ANSIC))

}
