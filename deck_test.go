package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// test deck length
	if len(d) != 54 {
		t.Errorf("Expected length of 54, but got %d", len(d))
	}

	// test deck order
	if d[0] != "A of Hearts" {
		t.Errorf("Expected 0th card to be 'A of Hearts', but got %s", d[0])
	}
	if d[18] != "5 of Diamonds" {
		t.Errorf("Expected 18th card to be '5 of Diamonds', but got %s", d[18])
	}
	if d[53] != "Monochrome Joker" {
		t.Errorf("Expected 53th card to be 'Monochrome Joker', but got %s", d[53])
	}
}
