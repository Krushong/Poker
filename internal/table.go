package internal

import "github.com/segmentio/ksuid"

type Table struct {
	ID          string
	CardOnTable []PlayingСard
	Bank        map[Player]int
	maxValue    int
}

func (t *Table) addCashInTableID(r *Room, money int, player Player) {

	for i, v := range *r {
		if player == v {
			(*r)[i].Cash -= money
			(*r)[i].LastBet = money
			t.Bank[player] += money
		}
	}

}

func NewTable() Table {
	return Table{
		ID:          ksuid.New().String(),
		CardOnTable: []PlayingСard{},
		Bank:        map[Player]int{},
	}
}

func (t *Table) GiveAllBankPlayer(player Player, room *Room) {

	for i, v := range *room {
		if v == player {
			(*room)[i].Cash = player.Cash + t.AllBank()
		}
	}

	t.Bank = make(map[Player]int)

}

func (t *Table) AllBank() int {

	var AllCash int

	for _, v := range t.Bank {
		AllCash = AllCash + v
	}

	return AllCash

}

func (t *Table) SplitBankBetweenAllWinningPlayers(players []Player, room *Room) {

	partBank := t.AllBank() / len(players)

	for i, v := range *room {
		for _, p := range players {
			if v == p {
				(*room)[i].Cash = p.Cash + partBank
			}
		}
	}

	t.Bank = make(map[Player]int)

}
