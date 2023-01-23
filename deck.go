package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

//creating a new deck
func newDeck() deck{
	cards :=deck{}

	cardSuits :=[]string {"Ace","Diamond", "Spades", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two","Three","Four","Five"}

	for _, suit :=range cardSuits{
		for _, values:= range cardValues{
			cards=append(cards,values+" of "+suit)
		}
	}

	return cards
}

//receiver function to print cards values
func (d deck) print(){
	for i,card :=range d{
		fmt.Println(i,card)
	}
}

//to get a deal from a deck
func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

//to convert deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//to save the deck to hard drive
func (d deck) saveToFile(fileName string) error{
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//to retrieve deckfrom file
func newDeckFromFile(filename string) deck{
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r :=rand.New(source)

	for i := range d{
		newPos := r.Intn(len(d)-1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}
