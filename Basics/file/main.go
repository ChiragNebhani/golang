package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("Exmaple.txt")
	if err != nil {
		fmt.Println("Error while creating file", err)
		return
	}
	defer file.Close()

	content := "Helo go!!!!!!!"
	_, error := io.WriteString(file, content)
	if error != nil {
		fmt.Println("Error while writing file : ", error)
		return
	}

	fmt.Println("File created succesfully")
}
