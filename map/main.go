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
	studentGrade["Gaurav"] = 90

	fmt.Println("Student Grade before deleting:", studentGrade)

	// Deleting Gaurav
	delete(studentGrade, "Gaurav")
	fmt.Println("Student Grade after deleting:", studentGrade)

	// Accessing marks of Garima
	fmt.Println("Marks of Garima is:", studentGrade["Garima"])

	// To check if a key exists
	grades, exists := studentGrade["Ronaldo"]
	fmt.Println("Grades of David is:", grades)
	fmt.Println("Does David exist in the map?", exists)

	for i, value := range studentGrade {
		fmt.Println("Grades of", i, "is:", value)
	}

	studentGrades := map[string] int {
		"Alice": 90,
		"Bob": 85,
		"Charlie": 95,
	}
	
	for i, value:=range studentGrades {
		fmt.Println("Grades of", i, "is:", value)
	}
}
