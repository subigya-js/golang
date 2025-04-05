package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World"
	parts := strings.Split(str, ",")

	fmt.Println(parts)

	str2 := "one two three four two two five"
	count := strings.Count(str2, "two")

	fmt.Println("The word 'two' appears", count, "times in the string.")

	str = "    Hello World     "
	str = strings.TrimSpace(str) // This removes leading and trailing whitespace

	fmt.Println("Trimmed string:", str)

	str = "Gaurav"
	str2 = "Subedi"

	joinedString := strings.Join([]string{str, str2}, " ")
	fmt.Println("Joined string:", joinedString)
}
