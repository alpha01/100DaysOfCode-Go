package main

import "fmt"

func main() {
	cards := newDeck()

	cards.wtf()
	fmt.Println("-----------------")

	hand, remainingCards := deal(cards, 5)
	fmt.Println("Hand:")
	hand.wtf()
	fmt.Println("-----------------")
	fmt.Println("Remaining:")
	remainingCards.wtf()
	fmt.Println("-----------------")
	cards.saveToFile("cards.txt")
}
