package Poker

import (
	"fmt"
	"os"
)

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

	playerBigBlaind := (*r)[(Dilerplayer+1)%len(*r)]
	playerSmallBlaind := (*r)[(Dilerplayer+2)%len(*r)]

	table.addCashInTableID(r, 50, playerBigBlaind)
	table.addCashInTableID(r, 25, playerSmallBlaind)
	table.maxValue = 50

	return table
}

func (r *Room) nextPlayerTrade() {
	//tyt novii raund
	//
	for _, p := range *r {
		if p.Diler {

		}
	}
}

// Добавить массив игроков в игре
func NextTradeRound(playersInGame *Room, table *Table) {

	var idAgressor int
	//var idPlayer int

	countPlayer := len(*playersInGame)

	for i := 0; i < countPlayer; i++ {
		if (*playersInGame)[i].Diler {
			idAgressor = returnCurentUser(i, countPlayer)
			break
		}
	}
	//Hod
	idCurrentUser := returnCurentUser(idAgressor, countPlayer)

	for true {
		for i := range *playersInGame {
			if idAgressor == 777 {
				break
			} else if len(*playersInGame) <= idCurrentUser {
				idCurrentUser = 0
			} else if idCurrentUser == i {
				move(playersInGame, table, i, &idAgressor)
				idCurrentUser++
			} else if i == idCurrentUser {
				idCurrentUser++
			}

		}
		if len(*playersInGame) == 1 {
			break
		} else if idAgressor == 777 {
			break
		}
	}
}

func returnCurentUser(i, countPlayer int) int {

	var idAgressor int

	if i+1 <= countPlayer {
		idAgressor = i + 1
	} else if i == countPlayer {
		idAgressor = 2
	}

	return idAgressor

}

//Статус 0 - Сбросить(пас) (Ставка -1 если хочешь сбросить) 1 -Пропустить 2 - Поднять
func move(playerInGame *Room, table *Table, idPlayer int, idAgressor *int) {

	answer := questionPlayer((*playerInGame)[idPlayer])

	switch {
	case answer == 0:
		playerInGame.fold(idPlayer)
		if *idAgressor == idPlayer {
			*idAgressor = 777
		}
	case answer == table.maxValue:
		cheak(table, playerInGame, (*playerInGame)[idPlayer])
		if *idAgressor == idPlayer {
			*idAgressor = 777
		}
	case answer > table.maxValue:
		raise(table, playerInGame, (*playerInGame)[idPlayer], answer)
		table.maxValue = answer
		*idAgressor = idPlayer
	}

}

func raise(table *Table, room *Room, player Player, money int) {

	table.addCashInTableID(room, money, player)

}

func cheak(table *Table, room *Room, player Player) {

	table.addCashInTableID(room, table.maxValue, player)

}

func (r *Room) fold(idPlayer int) {

	*r = append((*r)[:idPlayer], (*r)[idPlayer+1:]...)

}

func (playerInGame Room) LastPlayerInGame() bool {

	return len(playerInGame) == 1

}

func questionPlayer(player Player) int {

	var answer int

	println("Выберите ход")
	fmt.Fscan(os.Stdin, &answer)

	if player.Cash < answer {
		println("У вас нет столько фишек")
		questionPlayer(player)
	}

	if answer == -1 {
		return 0
	}

	Bet := player.LastBet + answer

	return Bet
}
