package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creating custom type Deck
//Which is same as slice of string

type deck []string

// Writing custom method
//1. create new deck (constructor)
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//2. print cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// 3. deal function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Helper method to convert deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// 4. saveToFile function
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// 5. read new deck from file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// 6. shuffle function
func (d deck) shuffle() {
	//Randomize the engine of pseudo random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
