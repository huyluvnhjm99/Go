package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.saveToFile("myPokemonCards.txt")
	// cards := newDeckFromFile("myPokemonCards.txt")
	// cards.print()
	// cards.shuffle()
	// cards.print()

	fmt.Println("Enter number to check: ")
	var input int
	fmt.Scan(&input)
	if checkEven(input) {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	var nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range nums {
		if checkEven(i) {
			fmt.Println(i, "is", "Even")
		} else {
			fmt.Println(i, "is", "Odd")
		}
	}
}
