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
