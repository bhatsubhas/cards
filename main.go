package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 3)
	fmt.Println("Cards in hand are:")
	hand.print()

	fmt.Println("Remaining cards in the deck are:")
	remainingCards.print()
}
