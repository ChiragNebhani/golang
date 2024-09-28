package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	part := strings.Split(data, ",")
	fmt.Println(part)

	data2 := "batman.superman.spiderman"
	part2 := strings.Split(data2, ".")
	fmt.Println(part2)

	str := "one two three two one"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str2 := "         Helo Go!    "
	trimmed := strings.TrimSpace(str2)
	fmt.Println(trimmed)

	firstName := "Chirag"
	lastName := "Nebhani"
	fullName := strings.Join([]string{firstName, lastName}, " ")
	fmt.Println(fullName)
}
