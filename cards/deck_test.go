package main

import "testing"

//这里的*testing.T 是test handler
func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expect 52 cards but got %v!", len(deck))
	} else if deck[0] != "Ace of Heart" {
		t.Errorf("first card should be a Ace of Heart if not shuffled!")
	}
}
