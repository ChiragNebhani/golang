package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Started error handling")
	ans, _ := divide(4, 2)
	fmt.Println(ans)
}
