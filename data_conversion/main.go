package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Numeric data conversion
	var integerNum int = 10
	var floatNum float64 = float64(integerNum)

	fmt.Println("Data type of integerNum", integerNum, "is:", reflect.TypeOf(integerNum))
	fmt.Println("Data type of floatNum",floatNum, "is:" ,reflect.TypeOf(floatNum))

	// String data conversion
	var number int = 10
	str := strconv.Itoa(number)
	fmt.Println("Data type of number", number, "is:", reflect.TypeOf(number))
	fmt.Println("Data type of str", str, "is:", reflect.TypeOf(str))

	strNum := "123"
	num, _ := strconv.Atoi(strNum)

	fmt.Println("Data type of strNum", strNum, "is:", reflect.TypeOf(strNum))
	fmt.Println("Data type of num", num, "is:", reflect.TypeOf(num))
}