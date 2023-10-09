package main

func main() {
	// cards := newDeck()
	
	// cards.print()

	// hand, remaindingDeck := deal(cards, 4)

	// hand.print()
	// remaindingDeck.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("the_deck.txt")

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	
	deckFromFile := newDeckFromFile("./the_deck.txt")

	deckFromFile.print()
	deckFromFile.shuffle()
	deckFromFile.print()
}