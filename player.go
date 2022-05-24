package Poker

import "github.com/segmentio/ksuid"

type Player struct {
	ID          string
	Hand        [2]PlayingСard
	InGame      bool
	Combination Combination
	Cash        int
	Diler       bool
	LastBet     int
}

type Combination struct {
	Combination    int8
	FirstMaxValue  int8
	SecondMaxValue int8
	Kicker         int8
}

func NewPlayer(ID string) Player {

	if ID == "" {
		player := Player{
			ID:          ksuid.New().String(),
			Hand:        [2]PlayingСard{},
			InGame:      true,
			Cash:        1000,
			Combination: Combination{},
			Diler:       false,
		}
		return player
	} else {
		player := Player{
			ID:          ID,
			Hand:        [2]PlayingСard{},
			InGame:      true,
			Cash:        1000,
			Combination: Combination{},
			Diler:       false,
		}
		return player
	}

}
