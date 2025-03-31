package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}

	return a / b, nil
}

func main() {
	fmt.Println("Error Handling in the functions.")

	divide, _ := divide(10, 0)
	fmt.Println("Division of 10 and 0 is: ", divide)
}
