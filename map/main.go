package main

import "fmt"

func main() {
	studentGrade := make(map[string]int)
	// This is somewhat similar to:
	// {
	// 	"name": <marks>,
	// }

	studentGrade["Gaurav"] = 90
	// {
	// 	"Gaurav": 90
	// }
	studentGrade["Garima"] = 85
	studentGrade["Subigya"] = 95

	fmt.Println(studentGrade)

	// Accessing marks of Garima
	fmt.Println("Marks of Garima is:", studentGrade["Garima"])
}
