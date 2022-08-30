package main

import (
	"os"
	"testing"
)

func TestGetNewDeck(t *testing.T) {
	deck := getNewDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if deck[0] != "Ace of Club" {
		t.Errorf("Expected first card in deck to be Ace of Club, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Two of Spade" {
		t.Errorf("Expected first card in deck to be Two of Spade, but got %v", deck[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := getNewDeck()

	deck.saveToFile("_decktesting")

	newDeck := newDeckFromFile("_decktesting")

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if len(newDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	os.Remove("_decktesting")

}
