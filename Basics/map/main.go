package main

import "fmt"

func main() {
	studentsGrades := make(map[string]int)

	studentsGrades["Chirag"] = 89
	studentsGrades["Lucky"] = 56
	studentsGrades["Yash"] = 55
	studentsGrades["Jodd"] = 78
	studentsGrades["Batman"] = 23

	fmt.Println("Marks of Chirag : ", studentsGrades["Chirag"])

	delete(studentsGrades, "Yash")
	fmt.Println("Marks of yash ", studentsGrades["Yash"])

	// checking if key exists
	Grades, Exists := studentsGrades["Chirag"]
	fmt.Println("Grades :", Grades)
	fmt.Println("Exists :", Exists)

	// unordered map
	for index, value := range studentsGrades {
		fmt.Printf("Index is %s and Value is %d\n", index, value)
	}
}
