package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	// Base returns the file name after the last slash.
	file := path.Base(os.Args[0])
	fmt.Println(file)

	// Dir returns the directory without the last file name.
	dir := path.Dir(os.Args[0])
	fmt.Println(dir)

	s, sep := "", ""
	for k, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Printf("Key: %v\t value: %v\n", k, arg)
	}
	fmt.Println(s)
}
