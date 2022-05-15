package internal

type Room []Player

func NewRoom() Room {
	return Room{}
}

func (r *Room) AddAllNewPlayerInRoom(PlayerID []string) {

	NumberPlayer := len(PlayerID)
	if NumberPlayer > 8 || NumberPlayer < 2 {
		return
	}

	for _, v := range PlayerID {
		player := NewPlayer(v)
		*r = append(*r, player)
	}

}

func (r *Room) AddNewPlayerInRoom(PlayerID string) {

	player := NewPlayer(PlayerID)
	*r = append(*r, player)

}

// переделать наименование
func (r *Room) MakeDiler(table Table) Table {

	Dilerplayer := 0

	(*r)[Dilerplayer].Diler = true

	table.addCashInTableID(r, 25, (Dilerplayer+1)%len(*r))
	table.addCashInTableID(r, 50, (Dilerplayer+2)%len(*r))

	return table
}

func (r *Room) nextPlayerTrade() {
	//tyt novii raund
	//
	for _, v := range *r {
		if v.Diler {

		}
	}
}

func (r *Room) nextTradeRound() {

	for _, p := range *r {
		if p.Diler {

		}
	}
}
