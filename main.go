package main

import "fmt"

func main() {
	cards := newDeck()
	s := cards.toString()
	fmt.Println(s)
}
