package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Creating a new type called deck, with the underlying type of slice of string type
type deck []string

// Receiver function
func (d deck) wtf() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Returns a value of type deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Receiver function that requires two arguments, deck type and int type.
// returns two deck type values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	t := []string(d) // convert to string
	return strings.Join(t, ",")
}

func (d deck) saveToFile(filename string) error {
	message := []byte(d.toString())

	return ioutil.WriteFile(filename, message, 0644)

}
