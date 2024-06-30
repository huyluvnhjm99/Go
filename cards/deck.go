package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"One", "Two", "Three", "Four"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toStrings() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toStrings()), 0666)
}

func newDeckFromFile(filname string) deck {
	byteSlice, err := os.ReadFile(filname)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
