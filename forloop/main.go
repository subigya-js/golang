package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Infinite loop:
	counter := 0

	for {
		counter++

		fmt.Println(counter)

		if counter == 5 {
			break
		}
	}
}
