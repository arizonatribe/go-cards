package main

func main() {
	cards := createDeck()
	cards.shuffle()
	cards.print()
}
