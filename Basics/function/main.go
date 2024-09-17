package main

import "fmt"

func simpleFunc() {
	fmt.Println("Simple Function")
}

func add(a, b int) int {
	return a + b
}

func mult(a int, b int) (result int) {
	return a * b
}

func main() {
	fmt.Println("Learning Functions in GO")
	simpleFunc()
	ans := add(3, 5)
	fmt.Println(ans)
	multAns := mult(4, 3)
	fmt.Println(multAns)
}
