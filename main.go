package main

func main() {
	//cards := newDeck()
	cards := loadFromFile("myDeck.txt", ",")
	cards.print()
}
