package main

import (
	"fmt"
)

func simpleFunction() {
	fmt.Println("This is a simple function")
}

func add(num1, num2 int) (result int) {
	result = num1 + num2

	return result
}

func main() {
	simpleFunction()

	sum := add(1, 2)
	fmt.Println("Sum of 1 and 2 is:", sum)
}
