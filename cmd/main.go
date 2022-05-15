package main

import (
	"fmt"
	"github.com/Krushong/Poker/internal"
)

func main() {

	table := internal.NewTable()

	PlayingDeck := internal.NewDeck()

	PlayingDeck.ShuffleDeck()

	room := internal.NewRoom()

	room.AddNewPlayerInRoom("")
	room.AddNewPlayerInRoom("")
	room.AddNewPlayerInRoom("")
	room.AddNewPlayerInRoom("")

	table = room.MakeDiler(table)

	PlayingDeck.GiveCardAllPlayerInRoom(room)

	//Стадия игры

	//стадия торгов
	PlayingDeck.Flop(&table)
	//стадия торгов
	PlayingDeck.TurnRiver(&table)
	//стадия торгов
	PlayingDeck.TurnRiver(&table)
	//стадия торгов

	fmt.Println(table.CardOnTable)

	winingPlayers := internal.WinChekerPlayers(room, table)

	if len(winingPlayers) == 1 {
		fmt.Println(winingPlayers[0].Hand)
		fmt.Println(winingPlayers[0].Combination)
	} else {
		for _, v := range winingPlayers {
			fmt.Println(winingPlayers[0].Hand)
			fmt.Println(v.Combination)
		}
	}
}
