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
		t.Errorf("Expected first element to be \"Ace of Spades\", but got \"%v\"", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last element to be \"King of Clubs\", but got \"%v\"", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	l := newDeckFromFile("_decktesting")

	if len(l) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(l))
	}

	os.Remove("_decktesting")

}
