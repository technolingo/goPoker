package main

import "fmt"

func main() {
	cards := newDeck()
	hand, deck := deal(cards, 5)
	hand.print()
	fmt.Println("========")
	deck.print()
}
