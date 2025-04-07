package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web requests in Go")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer res.Body.Close() // Close the response body when done

	fmt.Printf("Type or response is: %T\n", res)

	// Read the response data
	data, err := io.ReadAll(res.Body)
	// fmt.Printf("Type of data is: %T\n", data)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response data:", string(data))
}
