package Poker

import (
	"strconv"
)

// PlayingСard
//value 			- Card number
// valuePerfomans 	- in text 11 - J 12 - Q 13 - K 14 - A
// suit 			- Card suit 0 - Hearts 1 - Spades 2 - Diamonds 3 -Clubs
type PlayingСard struct {
	value          int8
	valuePerfomans string
	suit           int8
}

func FillCard(value, suit int8) PlayingСard {

	card := PlayingСard{
		value:          value,
		suit:           suit,
		valuePerfomans: returnValuePerfomans(value),
	}

	return card
}

func returnValuePerfomans(number int8) (valuePerfomans string) {

	switch number {
	case 11:
		valuePerfomans = "J"
	case 12:
		valuePerfomans = "Q"
	case 13:
		valuePerfomans = "K"
	case 14:
		valuePerfomans = "A"
	default:
		valuePerfomans = strconv.Itoa(int(number))
	}

	return

}
