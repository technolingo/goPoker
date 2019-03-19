package main

import (
	"os"
	"testing"
)

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

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	// delete previously created test file if any
	os.Remove("_decktesting")

	// create a shuffled deck
	saved := newDeck()
	saved.shuffle()

	// save the deck to file
	err := saved.saveToFile("_decktesting", ",")
	if err != nil {
		t.Error(err)
	}

	// load the deck from file
	loaded := loadFromFile("_decktesting", ",")

	// compare the loaded deck from the one created above
	if len(saved) != len(loaded) {
		t.Errorf("Expected len(saved) == len(loaded), but %d != %d", len(saved), len(loaded))
	}
	for i, s := range saved {
		if s != loaded[i] {
			t.Errorf("Expected saved[%d] == loaded[%d], but %s != %s", i, i, s, loaded[i])
		}
	}

	// delete the test file created above
	os.Remove("_decktesting")
}
