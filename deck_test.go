package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	cards := createDeck()

	if len(cards) != 52 {
		t.Errorf("Decks should have 52 cards, instead found %d", len(cards))
	}

	if cards[0] != "Ace of Clubs" {
		t.Errorf("First card in the deck should be the 'Ace of Clubs', but instead found %s", cards[0])
	}

	if cards[len(cards)-1] != "King of Spades" {
		t.Errorf("Last card in the deck should be the 'King of Spades', but instead found %s", cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	originDeck := createDeck()
	originDeck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Loaded a deck from a file that has %d cards instead of 52", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
