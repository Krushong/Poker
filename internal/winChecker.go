package internal

func WinChekerPlayer(room Room, table Table) {

	for i := range room {
		if room[i].InGame {
			room[i].Combination = cheackCombinationPlayer(room[i].Hand, table.CardOnTable)
		}
	}
}

func cheackCombinationPlayer(hand [2]PlayingСard, CardOnTable []PlayingСard) int8 {

	//Всего есть 10 комбинация
	//1 Старшая карта 2 Пара 3 Две пары 4 Сет 5 Стрит
	//6 Флеш 7 ФуллХаус 8 КАРЭ 9 Стрит Флэш 10 ФлешРояль

	allCard := CardOnTable

	for i := range hand {
		allCard = append(allCard, hand[i])
	}

	if royalFlushCheak(allCard) {
		return 10
	} else if straightFlushCheak(allCard) {

	}

	return 10
}

func royalFlushCheak(allCard []PlayingСard) bool {

	//if cheakSameSuits(allCard) {

	allValue := make([]int, 0)

	for i := range allCard {
		allValue = append(allValue, int(allCard[i].value))
	}

	fiveBestCard := findScoreFiveBestCard(allCard)

	score := int8(0)

	for _, v := range fiveBestCard {
		score = score + v.value
	}

	if score == int8(60) {
		return true
	} else {
		return false
	}

}

func FindDominantSuit(allCard []PlayingСard) int8 {

	allSuits := make([]int, 0)

	for i := range allCard {
		allSuits = append(allSuits, int(allCard[i].suit))
	}

	Hearts := findAllRepetitions(allSuits, 0)
	Spades := findAllRepetitions(allSuits, 1)
	Diamonds := findAllRepetitions(allSuits, 2)
	Clubs := findAllRepetitions(allSuits, 3)

	if Hearts > 4 {
		return 0
	} else if Spades > 4 {
		return 1
	} else if Diamonds > 4 {
		return 2
	} else if Clubs > 4 {
		return 3
	} else {
		return 4
	}
}

func findAllRepetitions(a []int, x int) int {

	allRepetitions := 0

	for _, n := range a {
		if x == n {
			allRepetitions++
		}
	}

	return allRepetitions

}

func findScoreFiveBestCard(allCard []PlayingСard) []PlayingСard {

	allCard = removeSmallestScore(allCard)
	if len(allCard) == 6 {
		allCard = removeSmallestScore(allCard)
	}
	return allCard
}

func removeSmallestScore(allCard []PlayingСard) []PlayingСard {

	smallestCard := 10

	allCard = removeAllOtherSuits(allCard)

	if len(allCard) == 5 {
		return allCard
	}

	for i, v := range allCard {
		s := -1
		for _, n := range allCard {
			if v.value <= n.value {
				s++
			}
		}

		if s == len(allCard)-1 {
			smallestCard = i
			break
		}
	}

	allCard = append(allCard[:smallestCard], allCard[smallestCard+1:]...)

	return allCard

}

func removeAllOtherSuits(allCard []PlayingСard) []PlayingСard {

	dominantSuit := FindDominantSuit(allCard)
	if dominantSuit == 4 {
		return allCard
	}

	for i, k := range allCard {
		if dominantSuit != k.suit {
			allCard = append(allCard[:i], allCard[i+1:]...)
			allCard = removeAllOtherSuits(allCard)
			return allCard
		}
	}

	return allCard
}

func straightFlushCheak(allCard []PlayingСard) bool {
	if allCard == nil {
	}
	return true
}
