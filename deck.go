package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{"Five of Diamonds",
		"Ace of Diamond",
		"Six of Heart",
		"Ace of Heart",
		"Queen of Heart"}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
