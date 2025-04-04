package main

import "fmt"

func modifyValue(modify *string) {
	*modify = "Hello World"
	fmt.Println("Value of num during modifying:", *modify)

	// Call by value:
	// fmt.Println("Value of num2:", num2)
	// fmt.Println("Pointer to num2:", ptr2)
}

func main() {
	var num string = "2"

	var ptr *string = &num

	num2 := "Hello"
	ptr2 := &num2
	fmt.Println("Value of num2:", num2)
	fmt.Println("Pointer to num2:", ptr2)

	fmt.Println("Address of num:", ptr)

	// Call by reference:
	fmt.Println("Value of num before modifying:", num)
	modifyValue(ptr)
	fmt.Println("Value of num after modifying:", num)
}
