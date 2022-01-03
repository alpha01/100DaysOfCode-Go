package main

import "fmt"

var a int

// creating a custom type
type hotdog int

var b hotdog

func main() {
	fmt.Printf("a is %T\n", a)
	b = 43
	fmt.Printf("b is %T\n", b)

	// cast b to int type
	a = int(b)
	fmt.Printf("a is %T\n", a)
}
