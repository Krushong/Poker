package main

import (
	"fmt"
	"github.com/Krushong/Poker"
)

func main() {

	table := Poker.NewTable()

	PlayingDeck := Poker.NewDeck()

	PlayingDeck.ShuffleDeck()

	room := Poker.NewRoom()

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

	winingPlayers := Poker.WinChekerPlayers(room, table)

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

func TradeRound(playerInGame, room *Poker.Room, table Poker.Table) {

	Poker.NextTradeRound(playerInGame, &table)
	if playerInGame.LastPlayerInGame() {
		table.GiveAllBankPlayer((*playerInGame)[0], room)
		return
	}

}
