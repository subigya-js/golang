package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter a number between 1 and 7: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Read input from user
	input = strings.TrimSpace(input)    // Remove newline character

	// Convert string into integer
	dayInput, _ := strconv.Atoi(input)

	switch dayInput {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}
}
