package main

func main() {

	cards := getNewDeck()

	log("New Deck")
	cards.printDeck()

	cardsInHand, remainingCards := deal(cards, 5)

	log("Cards In Hand")
	cardsInHand.printDeck()

	log("Remaining Cards")
	remainingCards.printDeck()

	log(cards.toString())
	cards.saveToFile("deck1.deck")

	newDeckFromFile("deck1.deck").printDeck()
	log(newDeckFromFile("deck1.deck").toString())

	cards.shuffle()
	cards.printDeck()

}
