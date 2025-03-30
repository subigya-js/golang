package main

import (
	"fmt"
	"mylearning/myutil"
	"reflect"
)

func PublicFunction() {
	fmt.Println("Public function")
}

func privateFunction() {
	fmt.Println("Private function")
}

func main() {
	fmt.Println("Hello world")

	myutil.PrintMessage("Hello from myutil package")

	// Variables

	// String
	var name string = "Gaurav"
	var lastname = "Subedi"
	fmt.Println(name)
	fmt.Println(lastname)

	fmt.Println("Type of name is: ", reflect.TypeOf(name))
	fmt.Println("Type of lastname is: ", reflect.TypeOf(lastname))

	// Integer
	var age = 23
	fmt.Println(age)
	var year int = 2002
	fmt.Println(year)
	fmt.Println("Type of age is: ", reflect.TypeOf(age))
	fmt.Println("Type of year is: ", reflect.TypeOf(year))

	// Float
	var cgpa = 8.77
	fmt.Println(cgpa)
	var percentage float64 = 87.7
	fmt.Println(percentage)

	// Boolean
	var isStudent = true
	fmt.Println(isStudent)
	var isWorking bool = true
	fmt.Println(isWorking)

	// Short Declaration
	fullName := "Gaurav Subedi"
	fmt.Println(fullName)
	fmt.Println("Type of fullName is: ", reflect.TypeOf(fullName))

	fullName = "Subigya Subedi"

	const PI = 3.14
	fmt.Println(PI)
	// PI = 3.14159 // This will give an error because PI is a constant

	var Public = "Public variable"
	var private = "private variable"
	fmt.Println("Public variable: ", Public)
	fmt.Println("Private variable: ", private)

	PublicFunction()
	privateFunction()

	fmt.Println("Public variable from myutil package: ", myutil.PublicVariable)
}
