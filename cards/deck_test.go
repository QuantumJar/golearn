package main

import (
	"os"
	"testing"
)

// 这里的*testing.T 是test handler
func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expect 52 cards but got %v!", len(deck))
	} else if deck[0] != "Ace of Heart" {
		t.Errorf("first card should be a Ace of Heart if not shuffled!")
	}
}

// 测试时创建一个临时文件_decktesting
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()

	deck.saveToFile("_deckTesting")

	newDeckFromFile := newDeckFromFile("_deckTesting")

	if len(newDeckFromFile) != 52 {
		t.Errorf("Expect deck length of 52 but found %v", len(newDeckFromFile))
	}
	os.Remove("_deckTesting")

}
