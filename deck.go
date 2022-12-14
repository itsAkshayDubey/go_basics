package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func getNewDeck() deck {

	cards := deck{}
	cardsSuites := []string{"Club", "Diamond", "Heart", "Spade"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}

	for _, suite := range cardsSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	cardsInHand := d[:handSize]
	remainingCards := d[handSize:]

	return cardsInHand, remainingCards
}

func log(s string) {
	fmt.Println(">>>>>> " + s + " <<<<<<")
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		log("Error occured while loading deck from file. Returning default deck.")
		return getNewDeck()
	}
	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
