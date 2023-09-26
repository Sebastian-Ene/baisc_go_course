package main

import (
	"fmt"
	"math/rand"
)

// Create new type

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}
	return cards
}

func (d deck) print() {
	fmt.Println("--- Printing deck ---")
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(numbOfCards int) (deck, deck) {
	// Created this before the course. It's basicaly a shuffle & deal
	hand := deck{}
	n := 0
	for n < numbOfCards {
		randomIndice := rand.Intn(len(d))
		hand = append(hand, d[randomIndice])
		d = append(d[:randomIndice], d[randomIndice+1:]...)
		n += 1
	}

	return hand, d
}

func dealCourse(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
