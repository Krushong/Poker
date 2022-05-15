package internal

import "github.com/segmentio/ksuid"

type Table struct {
	ID          string
	CardOnTable []PlayingСard
	Bank        map[Player]int
}

func (t *Table) addCashInTableID(r *Room, money, numberPlayer int) {

	(*r)[numberPlayer].Cash -= money
	t.Bank[(*r)[numberPlayer]] += money

}

func NewTable() Table {
	return Table{
		ID:          ksuid.New().String(),
		CardOnTable: []PlayingСard{},
		Bank:        map[Player]int{},
	}
}
