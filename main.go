package main

// import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// card := newCard()
	// count := 7.0

	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards,"Five of Diamonds")
	

	cards := newDeckFromFile("my_cards")

	cards.print()

	cards.shuffleDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("my_cards")

	cards.print()

	// fmt.Println(count)
}
