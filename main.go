package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println("Before shuffling")
	cards.print()
	fmt.Println("After shuffling")
	cards.shuffle()
	cards.print()
}
