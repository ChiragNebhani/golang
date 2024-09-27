package main

import "fmt"

func main() {
	day := 3
	switch day {
	case 1, 3:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown Day")
	}
}
