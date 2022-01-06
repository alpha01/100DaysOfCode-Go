package main

import "fmt"

func main() {
	cards := []string{"testing1", "testing2"}

	fmt.Println(cards)

	cards = append(cards, "testing3")
	fmt.Println(cards)

	for key, value := range cards {
		fmt.Println(key, value)
	}
}
