package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserID int    `json:"user_id"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Response struct {
	Status  string
	Body    string
	Headers map[string][]string
}

func sendPostRequest(url string, jsonData []byte) (*Response, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Status:  resp.Status,
		Body:    string(body),
		Headers: resp.Header,
	}

	return response, nil

}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	postData := Post{
		UserID: 1,
		ID:     1,
		Title:  "Test Post",
		Body:   "This is a test post.",
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	response, err := sendPostRequest(url, jsonData)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}

	// Print the response
	fmt.Println("Status:", response.Status)
	fmt.Println("Response Body:", response.Body)
}
