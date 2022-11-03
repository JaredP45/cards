package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for idx, card := range cards {
		fmt.Println(idx, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
