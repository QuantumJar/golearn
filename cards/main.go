package main

func main() {

	// d := newDeck()

	// hand, remaningDeck := deal(d, 3)

	// hand.print()
	// remaningDeck.print()
	// d.saveToFile()
	myDeck := newDeckFromFile("myDeck")
	// myDeck.print()

	myDeck.shuffle()

	myDeck.print()

}
