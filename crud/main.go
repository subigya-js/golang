package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", res.Status)
		return
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Response:", string(data))

	// We won't use the above approach to fetch the data
	// instead we will to the approach using structure

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Todo:", todo)
	fmt.Println("UserId:", todo.UserId)
}

func performPostRequest() {
	const myURL = "https://jsonplaceholder.typicode.com/todos"

	todo := Todo{
		UserId:    1,
		Id:        23,
		Title:     "Complete Go",
		Completed: true,
	}

	// Now convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the JSON data to a string
	jsonString := string(jsonData)

	// Convert the JSON string data into reader
	jsonReader := strings.NewReader(jsonString)

	// Now send the POST request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error while sending POST request:", err)
	}

	fmt.Println("Res", res)

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Post data:", string(data))

	// Now we have to close the resources of response
	defer res.Body.Close()
}

func performUpdateRequest() {
	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	todo := Todo{
		UserId:    2378,
		Id:        23,
		Title:     "Complete Go",
		Completed: true,
	}

	// Now convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the JSON data to a string
	jsonString := string(jsonData)

	// Convert the JSON string data into reader
	jsonReader := strings.NewReader(jsonString)

	// Create PUT request
	req, err := http.NewRequest("PUT", myURL, jsonReader)
	if err != nil {
		fmt.Println("PUT Request Error:", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send the request.
	client := http.Client{}
	res, _ := client.Do(req)

	readResonse, _ := io.ReadAll(res.Body)
	fmt.Println("Updated data:", string(readResonse))

	defer res.Body.Close()
}

func performDeleteRequest() {
	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	// Create Delete Request
	req, err := http.NewRequest("DELETE", myURL, nil)
	if err != nil {
		fmt.Println("PUT Request Error:", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send the request.
	client := http.Client{}
	res, _ := client.Do(req)

	fmt.Println("Deleted data status code:", res.Status)

	defer res.Body.Close()
}

func main() {
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}
