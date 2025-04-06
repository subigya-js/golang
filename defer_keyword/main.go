package main

import "fmt"

func main() {
		fmt.Println("Number 1")
		defer fmt.Println("Number 2")
		defer fmt.Println("Number 3")
		fmt.Println("Number 4")
}
