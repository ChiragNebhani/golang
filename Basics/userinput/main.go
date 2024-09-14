package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello,", name)

	// full name
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello,", name)
}
