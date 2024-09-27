package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 3, 4, 5, 6, 6, 7, 8, 9)
	fmt.Println("Number : ", numbers)
	fmt.Printf("Numbers has data type of %T\n", numbers)
	fmt.Println("Length of numbers : ", len(numbers))

	name := []string{}
	fmt.Println(name)

	numbers2 := make([]int, 3, 5)

	numbers2 = append(numbers2, 3)
	numbers2 = append(numbers2, 4)
	numbers2 = append(numbers2, 5)

	fmt.Println("SLice : ", numbers2)
	fmt.Println("Length :", len(numbers2))
	fmt.Println("Capacity :", cap(numbers2))
}
