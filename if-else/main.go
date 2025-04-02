package main

import "fmt"

func main() {
	x := 10
	y := 20

	if x > y {
		fmt.Println(x, "is greater than", y)
	} else {
		fmt.Println(y, "is greater than", x)
	}
}
