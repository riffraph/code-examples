package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type PlayingCard struct {
	Suit string
	Rank string
}

func NewPlayingCard(suit string, card string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: card}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

type Deck struct {
	cards []interface{}
}

func NewPlayingCardDeck() *Deck {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

func (d *Deck) AddCard(card interface{}) {
	d.cards = append(d.cards, card)
}

func (d *Deck) RandomCard() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

func main() {
	deck := NewPlayingCardDeck()

	fmt.Printf("--- drawing playing card ---\n")
	card := deck.RandomCard()
	fmt.Printf("drew card: %s\n", card)

	playingCard, ok := card.(*PlayingCard)
	if !ok {
		fmt.Printf("card received wasn't a playing card!")
		os.Exit(1)
	}
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)
}
