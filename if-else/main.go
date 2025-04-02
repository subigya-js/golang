package main

import "fmt"

func main() {
	x := 10
	y := 20

	if x > y {
		fmt.Println(x, "is greater than", y)
	} else if x == y {
		fmt.Println(x, "is equal to", y)
	} else {
		fmt.Println(y, "is greater than", x)
	}
}
