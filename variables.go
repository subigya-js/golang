package main

import "fmt"

// DemonstrateVariables shows examples of variable declarations and usage
func DemonstrateVariables() {
    // Declaring variables
    var a string = "initial"
    var b, c int = 1, 2
    var d = true

    // Short declaration
    e := 2.71828

    // Printing variables
    fmt.Println("Variable examples:")
    fmt.Println("a:", a)
    fmt.Println("b:", b)
    fmt.Println("c:", c)
    fmt.Println("d:", d)
    fmt.Println("e:", e)
}
