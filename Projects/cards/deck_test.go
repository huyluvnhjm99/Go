package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}

	if d[0] != "One of One" {
		t.Errorf("Expected first card is 'One of One', but got '%v'", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	deck := newDeck()
	deck.saveToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")
}

// In the upcoming lecture, we will be running some tests using Errorf. Omitting a formatting directive will now cause the tests to fail, so, we will need to add these right away.
// Note - this directive was originally added at the end of the lecture.

// When running your tests starting around the 8:00 timestamp, change the following:
// t.Errorf("Expected deck length of 16, but got", len(d))
// to this:
// t.Errorf("Expected deck length of 16, but got %v", len(d))
