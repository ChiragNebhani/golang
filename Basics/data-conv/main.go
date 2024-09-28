package main

import (
	"fmt"
	"strconv"
)

func main() {
	var Num int = 42
	fmt.Println("Num is ", Num)
	fmt.Printf("Type of Num is %T\n", Num)

	var Data float64 = float64(Num)
	fmt.Println("Data is ", Data)
	fmt.Printf("Type of Data is %T\n", Data)

	str := strconv.Itoa(Num)
	fmt.Println("str is ", str)
	fmt.Printf("Type of str is %T\n", str)

	number_string := "1234"
	number_integer, _ := strconv.Atoi(number_string)
	fmt.Println("number_integer is ", number_integer)
	fmt.Printf("Type of number_integer is %T\n", number_integer)

}
