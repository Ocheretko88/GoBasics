package main

import (
	"fmt"
	"time"
)

func main() {
	// If-Else
	t := time.Now()
	if t.Hour() < 12 {
		fmt.Println("It's before noon")
	} else {
		fmt.Println("It's after noon")
	}

	// For Loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Switch Case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
