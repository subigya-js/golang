package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json: "userId"`
	Id        int    `json: "id"`
	Title     string `json "title"`
	Completed bool   `json "completed"`
}

func main() {
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
