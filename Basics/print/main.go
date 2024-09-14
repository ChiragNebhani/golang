package main

import "fmt"

func main() {
	age := 21
	name := "Chirag Nebhani"
	position := "Golang Developer"
	height := 185
	weight := 75

	fmt.Println("Age : ", age, ", Name : ", name, ", Position : ", position)

	fmt.Printf("Height : %d\n", height)
	fmt.Printf("Weight : %d\n", weight)
	fmt.Printf("Type of name is %T\n", name)
}
