package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsStudent bool   `json:"is_student"`
}

func main() {
	person := Person{
		Name:      "Gaurav",
		Age:       23,
		IsStudent: true,
	}

	fmt.Println("Person Data:", person)

	// Convert to JSON (Encoding / Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	// Decoding / Unmarshalling
	var personData Person
	err = json.Unmarshal(jsonData, &personData)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Decoded Person Data:", personData)
}
