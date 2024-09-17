package main

import "fmt"

func main() {
	fmt.Println("Array in GO")

	var arr [5]string
	arr[0] = "Chirag"
	arr[2] = "Jodd"
	arr[4] = "Batman"

	fmt.Println("The names in array are ", arr)

	var numbers = [8]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("The length of number array is ", len(numbers))

	fmt.Println("Value of name at 2nd index is ", arr[2])
}
