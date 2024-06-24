package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	fmt.Println("Cards in hand: ")
	hand.print()
	fmt.Println("Cards remaining: ")
	remainingCards.print()
}
