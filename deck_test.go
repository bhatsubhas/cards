package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ten of Diamonds" {
		t.Errorf("Expected last card of Ten of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("__decktesting")
	d := newDeck()
	d.saveToFile("__decktesting")

	loadedDeck := newDeckFromFile("__decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck got %v", len(loadedDeck))
	}

	os.Remove("__decktesting")
}
