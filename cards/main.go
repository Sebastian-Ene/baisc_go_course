package main

func main() {
	cards := newDeck()

	// cards.print()

	hand1, cards := cards.deal(4)

	hand1.print()
	cards.print()

	hand2, cards := dealCourse(cards, 2)

	hand2.print()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
