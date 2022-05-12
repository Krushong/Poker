package internal

import "github.com/segmentio/ksuid"

type Player struct {
	ID          string
	Hand        [2]PlayingСard
	InGame      bool
	Score       int8
	Combination int8
}

func NewPlayer(ID string) Player {

	if ID == "" {
		player := Player{
			ID:          ksuid.New().String(),
			Hand:        [2]PlayingСard{},
			InGame:      true,
			Score:       0,
			Combination: 0,
		}
		return player
	} else {
		player := Player{
			ID:          ID,
			Hand:        [2]PlayingСard{},
			InGame:      true,
			Score:       0,
			Combination: 0,
		}
		return player
	}

}
