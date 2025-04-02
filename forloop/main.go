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

		if counter == 10 {
			break
		}
	}

	// Ranger keyword:
	for index, value:=range numbers {
		fmt.Println("The number in index",index, "is:", value)
	}

	// For string
	text := "Hello world"

	for index, value:=range text {
		fmt.Println("The character in index",index, "is:", string(value))
	}
}
