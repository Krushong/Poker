package Poker

import (
	"math/rand"
	"time"
)

type Deck []PlayingСard

func NewDeck() Deck {

	var value int8
	var suit int8
	var NewDeck Deck

	value = 2

	for i := 0; i < 52; i++ {

		NewDeck = append(NewDeck, FillCard(value, suit))

		value++
		if value == 15 {
			value = 2
			suit++
		}
	}

	return NewDeck

}

func (d Deck) ShuffleDeck() {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d),
		func(i, j int) { d[i], d[j] = d[j], d[i] })

}

func (d *Deck) cutCard(number int8) {
	*d = (*d)[number:]
}

func (d *Deck) Flop(table *Table) {
	d.cutCard(1)
	for i := 0; i < 3; i++ {
		table.CardOnTable = append(table.CardOnTable, (*d)[i])
	}
	d.cutCard(3)
}

func (d *Deck) TurnRiver(table *Table) {
	d.cutCard(1)
	table.CardOnTable = append(table.CardOnTable, (*d)[0])
	d.cutCard(1)
}

func (d *Deck) GiveCardAllPlayerInRoom(room Room) {

	CircleOfCards(0, room, d)
	CircleOfCards(1, room, d)

}

func CircleOfCards(circle int8, room Room, d *Deck) {

	for i := range room {
		room[i].Hand[circle] = (*d)[0]
		d.cutCard(1)
	}

}
