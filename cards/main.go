package main

func main() {
	// create a new deck
	//cards := newDeck()

	// save to file
	//cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)
}
