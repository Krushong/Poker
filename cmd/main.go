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

	PlayingDeck.GiveCardAllPlayerInRoom(room)

	//Стадия игры

	//стадия торгов
	playerInGame := room
	table = playerInGame.MakeDiler(table)

	TradeRound(&playerInGame, &room, table)
	PlayingDeck.Flop(&table)
	TradeRound(&playerInGame, &room, table)
	PlayingDeck.TurnRiver(&table)
	TradeRound(&playerInGame, &room, table)
	PlayingDeck.TurnRiver(&table)
	TradeRound(&playerInGame, &room, table)

	fmt.Println(table.CardOnTable)

	winingPlayers := internal.WinChekerPlayers(room, table)

	if len(winingPlayers) == 1 {
		table.GiveAllBankPlayer(winingPlayers[0], &room)
	} else if len(winingPlayers) > 1 {
		table.SplitBankBetweenAllWinningPlayers(winingPlayers, &room)
	}

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

func TradeRound(playerInGame, room *internal.Room, table internal.Table) {

	internal.NextTradeRound(playerInGame, &table)
	if playerInGame.LastPlayerInGame() {
		table.GiveAllBankPlayer((*playerInGame)[0], room)
		return
	}

}
