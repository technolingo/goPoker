package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type deck []string

// newDeck func creates a full & ordered deck of cards.
func newDeck() deck {
	suits := [4]string{"Hearts", "Spades", "Diamonds", "Clubs"}
	// declare a ranks slice with initial value 'A'
	ranks := []string{"A"}
	// regular ranks
	for i := 2; i <= 10; i++ {
		ranks = append(ranks, strconv.Itoa(i))
	}
	// add other irregular ranks (could also add these along with Ace)
	ranks = append(ranks, []string{"J", "Q", "K"}...)

	// declare a cards deck to be returned
	var cards deck

	// combine ranks with suits
	for _, r := range ranks {
		for _, s := range suits {
			cards = append(cards, r+" of "+s)
		}
	}
	// add coloured & monochrome jokers
	cards = append(cards, "Coloured Joker")
	cards = append(cards, "Monochrome Joker")
	return cards
}

// print is a method of the deck type.
// It prints out all the cards in the deck.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal func takes a deck of cards & a deck size.
// And it returns a hand of cards & the remaining deck.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString is a method of the deck type.
// It converts the deck into one long string with a given delimiter.
func (d deck) toString(delimiter string) string {
	return strings.Join(d, delimiter)
}

// saveToFile is a method of the deck type.
// It saves the deck to a file with a given filepath and delimiter.
func (d deck) saveToFile(filepath string, delimiter string) error {
	return ioutil.WriteFile(filepath, []byte(d.toString(delimiter)), 0666)
}

// loadFromFile func loads the deck from a given file.
func loadFromFile(filepath string, delimiter string) deck {
	bs, err := ioutil.ReadFile(filepath)
	// error handling
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), delimiter))
}
