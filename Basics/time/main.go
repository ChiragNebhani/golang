package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Time is ", currentTime)
	fmt.Printf("Type of currentTime is %T", currentTime)

	formatted := currentTime.Format("02-01-2006")
	fmt.Println("Formatted time ", formatted)
}
