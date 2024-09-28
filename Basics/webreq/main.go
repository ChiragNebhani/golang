package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response")
		return
	}
	defer res.Body.Close()

	// read the response body
	data, error := ioutil.ReadAll(res.Body)
	if error != nil {
		fmt.Println("Erorr reading response ", error)
		return
	}

	fmt.Println("response : ", string(data))
}
