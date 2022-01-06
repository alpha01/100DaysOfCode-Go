package main

import "fmt"

func main() {
	t := testFunc()
	p := tonyTest()
	fmt.Printf("The value of t is '%v'\n", t)
	fmt.Printf("The value of p is '%v'\n", p)
	printState()
}

func testFunc() string {
	return "Return value"
}

func tonyTest() int {
	return 100
}
