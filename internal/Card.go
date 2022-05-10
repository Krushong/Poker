package internal

import "strconv"

type PlayingСards struct {
	value          int8
	valuePerfomans string
	suit           int8
}

var deck [52]*PlayingСards

func NewDeck() *[52]*PlayingСards {

	var value int8
	var suit int8

	for _, v := range deck {

		value = 2
		suit = 0
		v.value = value
		v.valuePerfomans = returnValuePerfomans(value)
		v.suit = suit

		value++
		if value == 14 {
			value = 2
			suit++
		}
	}

	return &deck

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
