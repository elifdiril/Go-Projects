package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeckFromFile("my_cards")
	fmt.Println(deck)
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", deck[len(deck)-1])
	}

	expected := []string{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades"}
	for i, card := range deck[:5] {
		if card != expected[i] {
			t.Errorf("Expected %s, but got %s", expected[i], card)
		}
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}