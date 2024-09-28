package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learnig url")

	myurl := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("Type url : %T\n", myurl)

	parsed_url, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Can't parse url", err)
		return
	}
	fmt.Printf("Type url : %T", parsed_url)
}
