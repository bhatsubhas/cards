package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println("Converting to a string")
	fmt.Println(cards.toString())

	fmt.Println("Saving to a file")
	cards.saveToFile("mycards.txt")

	fmt.Println("Reading from a file")
	mycards := newDeckFromFile("mycards.txt")
	mycards.print()

	hand, remainingCards := deal(cards, 3)
	fmt.Println("Cards in hand are:")
	hand.print()

	fmt.Println("Remaining cards in the deck are:")
	remainingCards.print()
}
