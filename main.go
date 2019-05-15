package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// Append doesn't add to existing cards, it creates a new slice
	cards = append(cards, "Six of Spades")

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
