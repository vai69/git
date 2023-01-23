package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 25 {
		t.Errorf("Expected deck length to be 25 got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 25 {
		t.Errorf("Expected deck length to be 25 got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
