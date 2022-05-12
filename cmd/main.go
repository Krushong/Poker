package main

import (
	"github.com/Krushong/Poker/internal"
)

func main() {

	table := internal.NewTable()

	PlayingDeck := internal.NewDeck()

	PlayingDeck.ShuffleDeck()

	room := internal.NewRoom()

	room.AddNewPlayerInRoom("")

	PlayingDeck.GiveCardAllPlayerInRoom(room)

	//Стадия игры
	PlayingDeck.Flop(&table)
	PlayingDeck.TurnRiver(&table)
	PlayingDeck.TurnRiver(&table)

	internal.WinChekerPlayer(room, table)
}
