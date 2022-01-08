package main

import (
	"fmt"
	"path"
)

func main() {

	example := "/home/bird"

	// Base returns the file name after the last slash.
	file := path.Base(example)
	fmt.Println(file)

	// Dir returns the directory without the last file name.
	dir := path.Dir(example)
	fmt.Println(dir)
}
