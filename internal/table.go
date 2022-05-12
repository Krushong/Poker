package internal

import "github.com/segmentio/ksuid"

type Table struct {
	ID          string
	CardOnTable []PlayingСard
}

func NewTable() Table {
	return Table{
		ID:          ksuid.New().String(),
		CardOnTable: []PlayingСard{},
	}
}
