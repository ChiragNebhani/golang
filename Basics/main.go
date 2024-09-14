package main

import (
	"fmt"
	"myLearning/myutil"
)

func main() {
	fmt.Println("Learn Go Language by Hello World")
	myutil.PrintMessage("Hello World")

	// variables
	// var name string = "Chirag"
	// var version = "1.0"
	// fmt.Println(name)
	// fmt.Println(version)

	// var currency = 34567
	// fmt.Println("Currency is", currency)

	// var num float64 = 1.65
	// fmt.Println(num)

	// var istrue bool = true
	// fmt.Println(istrue)

	// var golangDeveloper string = "Chirag Nebhani"
	// fmt.Println(golangDeveloper)

	// const pi = 3.14
	// fmt.Println(pi)

	person := "Chirag"
	println(person)

	var Public = "exportable"
	var private = "not exportable"
	fmt.Println(Public, private)
}
