package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, world!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("sayHello function execution completed.")
}

func sayHi() {
	fmt.Println("Hi! :)")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("sayHi function execution completed.")
}

func main() {
	fmt.Println("Learning Goroutine")
	go sayHello()
	go sayHi()

	time.Sleep(2000 * time.Millisecond)
}
