package main

import "fmt"

func main() {
	cards := newDeck()

	// cards.print()

	hand1, cards := cards.deal(4)

	hand1.print()
	cards.print()

	hand2, cards := dealCourse(cards, 2)

	hand2.print()
	cards.print()

	newDeck := newDeck()
	fmt.Println(newDeck.toString())
	newDeck.saveToFile("something.txt")

	newDeck1 := newDeckFromFile("something.txt")
	newDeck1.print()

	newDeck1 = newDeck1.shuffle()
	newDeck1.print()

}

func newCard() string {
	return "Five of Diamonds"
}
