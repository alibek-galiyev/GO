package main

func main() {
	// cards := newDeck()
	// cards, err := readFile("my_cards")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(cards)

	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("Hand:")
	// hand.print()
	// fmt.Println("Remaining deck:")
	// remainingDeck.print()

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
