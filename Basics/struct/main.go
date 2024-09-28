package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// 1
	var person1 Person
	fmt.Println("Person1 : ", person1)
	person1.firstName = "Chirag"
	person1.lastName = "Nebhani"
	person1.age = 20
	fmt.Println("Person1 data : ", person1)

	// 2
	person2 := Person{
		firstName: "Bat",
		lastName:  "Man",
		age:       0,
	}
	fmt.Println("Person2 data :", person2)

	// 3 - new keyword
	var person3 = new(Person)
	person3.firstName = "New"
	person3.lastName = "Keyword"
	person3.age = 200

	fmt.Println("Person3 data : ", person3)
}
