package main

import (
	"fmt"
	"math/rand"
	"time"
)


type Deck struct {

}
var cards = make([] Card, 0)

func populateDeck()  {

	for i := 2; i < 15; i++ {//Ace is 14
		card := new(Card)
		card.face = i
		card.suit = "Hearts"
		cards = append(cards, *card)
	}
	for i := 2; i < 15; i++ {//Ace is 14
		card := new(Card)
		card.suit = "Diamonds"
		card.face = i
		cards = append(cards, *card)
	}
	for i := 2; i < 15; i++ {//Ace is 14
		card := new(Card)
		card.suit = "Clubs"
		card.face = i
		cards = append(cards, *card)
	}
	for i := 2; i < 15; i++ {//Ace is 14
		card := new(Card)
		card.suit = "Spades"
		card.face = i
		cards = append(cards, *card)
	}
}

func hitMe() Card{
	returnCard := cards[len(cards)-1]
	//printDeck()
	cards = cards[:len(cards)-1]
	return returnCard
}

func shuffleDeck()  {
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 4; j++ {
		for i := 0; i < 51; i++ {
			index := rand.Intn(10)

			var tmpC Card = cards[i]
			cards[i] = cards[index]
			cards[index] = tmpC

		}
	}
}

func printDeck()  {
	for _, card := range cards {
		fmt.Println(card)
	}
}

