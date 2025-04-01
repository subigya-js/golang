package main

import "fmt"

func main () {
	fmt.Println("Array in Go.")

	var name [5] string // Array with default values

	name[0] = "Gaurav"
	name[2] = "Subigya"

	var marks = [5] int { 10, 20, 30, 40, 50 } // Array with initialization
	
	fmt.Println(name)
	fmt.Println(marks)
	fmt.Println("Length of array marks is:", len(marks))

	var test[5] bool
	fmt.Println(test)
}