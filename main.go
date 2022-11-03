package main

import "fmt"

func main() {
	card := newCard()
	// can be initialized as var card string = "Ace of Spades"

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
