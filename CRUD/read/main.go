package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("hello world from web request")
	// url := "https://jsonplaceholder.typicode.com/todos/1"
	url2 := "https://jsonplaceholder.typicode.com/posts"

	// Indicating the start of the request
	fmt.Println("Starting data loading...")

	// Making the HTTP GET request
	response, err := http.Get(url2)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer response.Body.Close() //close the connection

	// Indicating that data is loading
	fmt.Println("Data is loading...")

	// Reading the response body
	bodyBytes, err := ioutil.ReadAll(response.Body) //read the data in json format
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Indicating that loading is complete
	fmt.Println("Finished loading data:")
	fmt.Println(string(bodyBytes))
}
