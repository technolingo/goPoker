package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// a custom type 'deck', a slice of strings
type deck []string

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

// func (receiver) func-name(args) return-type {}
// any var of type 'deck' can access this print method
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal a given number of cards from a given deck
// return a hand of cards & the remaining deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// a helper method that converts a given deck into
// one large string with a given delimiter
func (d deck) toString(delimiter string) string {
	return strings.Join(d, delimiter)
}

// save a given deck to a file
func (d deck) saveToFile(filepath string, delimiter string) error {
	return ioutil.WriteFile(filepath, []byte(d.toString(delimiter)), 0666)
}

// load a deck from a given file
func loadFromFile(filepath string, delimiter string) deck {
	bs, err := ioutil.ReadFile(filepath)
	// error handling
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), delimiter))
}
