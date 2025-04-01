package main

import "fmt"

func main() {
	// numbers := []int{10, 20, 30, 40, 50}
	// fmt.Println("Slice numbers :", numbers)

	// // Length of slice
	// fmt.Println("Length of slice number is: ", len(numbers))

	// // To add elements in slice
	// numbers = append(numbers, 60, 70)
	// fmt.Println("Slice numbers after appending 60 and 70: ", numbers)
	// fmt.Println("Length of new slice number is: ", len(numbers))

	// // Working with 0 length slice
	// name := [] string {}
	// name = append(name, "Gaurav")
	// fmt.Println(name)

	// Using make function
	// ages := make([]int, 3, 5)
	// ages = append(ages, 10)
	// ages = append(ages, 20)
	// ages = append(ages, 30)
	// fmt.Println("Slice ages: ", ages)
	// fmt.Println("Length of slice ages: ", len(ages))
	// fmt.Println("Capacity of slice ages: ", cap(ages))

	// Zero slice
	stock := make([] int, 0)
	fmt.Println("Slice stock: ", stock)
	fmt.Println("Length of slice stock: ", len(stock))
	fmt.Println("Capacity of slice stock: ", cap(stock))
}
