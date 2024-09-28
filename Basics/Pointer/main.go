package main

import "fmt"

func modifyByReference(num *int) {
	*num = *num * 20
}

func main() {
	// var num int
	// num = 2
	num := 2

	// var ptr *int
	// ptr = &num
	ptr := &num

	fmt.Println("Num contains ", num)
	fmt.Println("ptr contains ", ptr)
	fmt.Println("data in ptr ", *ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer value is nil")
	}

	val := 10
	modifyByReference(&val)
	fmt.Println("Value of val is ", val)
}
