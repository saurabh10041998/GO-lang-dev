package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected card length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card of Ace of Diamonds but got %v", d[0])

	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of deck to be 16 , got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
