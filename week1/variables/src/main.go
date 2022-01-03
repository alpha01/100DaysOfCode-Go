package main

import "fmt"

// variable is globally scoped
var a = 40

// assign a zero value global variable
var b int

// assign an empty string global variable
var tony string

var c bool

func main() {
	fmt.Println("testing")

	// declare and assign a value using short declaration operator
	x := 35

	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(tony)
	fmt.Println(c)
	c = true
	fmt.Println(c)
}
