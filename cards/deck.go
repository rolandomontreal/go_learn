package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck' - a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + suit)
		}
	}
	return cards;
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
  fmt.Println()
}

func deal(d deck, handSize int) (deck, deck) {
  return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
  return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
  content := d.toString()
  return os.WriteFile(filename, []byte(content), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, error := os.ReadFile(filename)

	if error != nil {
		fmt.Println("Error reading from file: ", error)
		os.Exit(1)
	}

	stringifiedDeck := string(bs)

	newDeck := strings.Split(stringifiedDeck, ",")

	return newDeck
}

func (d deck) shuffle() {
	for i := range d {
		swapIndex := rand.Intn(len(d))
		d[i], d[swapIndex] = d[swapIndex], d[i]
	}
}