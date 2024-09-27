package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Count ", i)
	}

	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++
		if counter == 3 {
			break
		}
	}

	// range key word
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Value %d & index %d\n", value, index)
	}
}
