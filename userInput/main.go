package main

import "fmt"

func main() {
	var name string

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Hello", name, "welcome to the world of Go programming!")
	// fmt.Printf("Hello %s welcome to the world of Go programming!\n", name)
}
