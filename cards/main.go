package main

func main() {
	cards := newDeck()
	//hand, remainingCards := deal(cards, 5)
	//deck := newDeckFromFile("my_card")
	cards.shuffle()
	cards.print()
}
