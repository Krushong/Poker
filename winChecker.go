package Poker

import "sort"

func WinChekerPlayers(room Room, table Table) []Player {

	for i := range room {
		if room[i].InGame {
			room[i].Combination.Combination, room[i].Combination.FirstMaxValue,
				room[i].Combination.SecondMaxValue, room[i].Combination.Kicker = cheackCombinationPlayer(room[i].Hand, table.CardOnTable)
		}
	}

	return winningPlayerInRoom(room)

}

func cheackCombinationPlayer(hand [2]PlayingСard, CardOnTable []PlayingСard) (int8, int8, int8, int8) {

	//Всего есть 10 комбинация
	//1 Старшая карта 2 Пара 3 Две пары 4 Сет 5 Стрит
	//6 Флеш 7 ФуллХаус 8 КАРЭ 9 Стрит Флэш 10 ФлешРояль
	//Кикер реализован не совсем правильно сравнение должно идти по 5 картам где кикер это старшая карта
	//Тоесть побеждает не тот игрок у кого карта старше а тот у кого 2 по старшенсу 3 по старшенству 4 по старшенсву
	// или 5 по старшенству если же нет таких карт тогда банк делится - Это справедливо для пары и старшей карты
	//Для сета (тройки) сравнивается старшая 4 и 5 карта если они равны ничья
	//Если 2 пары победитель определяется по кикеру 5 старшей карте на столе
	//Фулл хаус выигрывает по более старшей карте
	//Каре если оно собралось на столе выигрывает тот у кого более старшая 2 карта

	allCard := CardOnTable

	for i := range hand {
		allCard = append(allCard, hand[i])
	}

	kicker := highCard(hand[0].value, hand[1].value)

	if royalFlushCheak(allCard) {
		return 10, 0, 0, 0
	} else if ok, maxValue := straightFlushCheak(allCard); ok {
		return 9, maxValue, kicker, 0
	} else if ok, maxValue := fourOfAKindCheak(allCard); ok {
		return 8, maxValue, kicker, 0
	} else if ok, firstMaxValue, secondMaxValue := fullHouseCheak(allCard); ok {
		return 7, firstMaxValue, secondMaxValue, kicker
	} else if ok, firstMaxValue, secondMaxValue := flushCheak(allCard); ok {
		return 6, firstMaxValue, secondMaxValue, kicker
	} else if ok, maxValue := straightCheak(allCard); ok {
		return 5, maxValue, kicker, 0
	} else if ok, maxValue := threeOfAKindCheak(allCard); ok {
		return 4, maxValue, kicker, 0
	} else if ok, firstMaxValue, secondMaxValue := twoPairCheak(allCard); ok {
		return 3, firstMaxValue, secondMaxValue, kicker
	} else if ok, maxValue := pairCheak(allCard); ok {
		return 2, maxValue, kicker, 0
	} else {
		return 1, kicker, 0, 0
	}

}

func royalFlushCheak(allCard []PlayingСard) bool {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	fiveBestCard := findScoreFiveBestCard(copyAllCard)

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

func straightFlushCheak(allCard []PlayingСard) (bool, int8) {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	dominantSuit := FindDominantSuit(copyAllCard)
	if dominantSuit == 4 {
		return false, 0
	}

	copyAllCard = removeAllOtherSuits(copyAllCard)
	if len(copyAllCard) <= 4 {
		return false, 0
	} else {
		return findStreat(copyAllCard)
	}

}

func cheakCardInDeck(a []PlayingСard, x int8) bool {

	for _, n := range a {
		if x == n.value {
			return true
		}
	}
	return false

}

func sortCardInOrder(allCard []PlayingСard) []PlayingСard {

	sort.SliceStable(allCard, func(i, j int) bool {
		return allCard[i].value < allCard[j].value
	})

	return allCard

}

func fourOfAKindCheak(allCard []PlayingСard) (bool, int8) {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	if ok, value := RepeatingCard(copyAllCard, 4); ok {
		return true, value
	} else {
		return false, 0
	}

}

func RepeatingCard(allCard []PlayingСard, valueRepeating int8) (bool, int8) {

	var s int8
	var value int8

	for _, v := range allCard {

		s = 0

		for _, n := range allCard {
			if v.value == n.value {
				s++
			}
		}

		if s == valueRepeating {
			value = v.value
			break
		}

	}

	if s == valueRepeating {
		return true, value
	} else {
		return false, value
	}
}

func fullHouseCheak(allCard []PlayingСard) (bool, int8, int8) {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	if ok, firstMaxValue := RepeatingCard(copyAllCard, 3); ok {

		copyAllCard = removeAllFindedCard(copyAllCard, firstMaxValue)

		if ok, secondMaxValue := RepeatingCard(copyAllCard, 2); ok {
			return true, firstMaxValue, secondMaxValue
		} else {
			return false, 0, 0
		}
	} else {
		return false, 0, 0
	}

}

func removeAllFindedCard(allCard []PlayingСard, value int8) []PlayingСard {

	for i, v := range allCard {
		if v.value == value {
			allCard = append(allCard[:i], allCard[i+1:]...)
			allCard = removeAllFindedCard(allCard, value)
			return allCard
		}
	}

	return allCard

}

func flushCheak(allCard []PlayingСard) (bool, int8, int8) {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	dominantSuit := FindDominantSuit(copyAllCard)
	if dominantSuit == 4 {
		return false, 0, 0
	}
	copyAllCard = removeAllOtherSuits(copyAllCard)

	if len(copyAllCard) >= 5 {
		maxValue := findMaxValue(copyAllCard)

		return true, maxValue, 0
	}

	return false, 0, 0
}

func findMaxValue(allCard []PlayingСard) int8 {

	var maxValue int8

	for _, v := range allCard {
		for _, n := range allCard {
			if v.value > n.value {
				maxValue = v.value
			}
		}
	}

	return maxValue
}

func straightCheak(allCard []PlayingСard) (bool, int8) {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	return findStreat(copyAllCard)

}

func findStreat(allCard []PlayingСard) (bool, int8) {

	allCard = sortCardInOrder(allCard)

	var lastValue int8
	var s int8
	var maxValue int8

	A := cheakCardInDeck(allCard, 14)
	for _, v := range allCard {

		if lastValue+1 == v.value || (A && v.value == 2) {
			s++
			if maxValue < v.value || (A && v.value == 2) {
				maxValue = v.value
			}
		}

		lastValue = v.value

	}

	if s >= 4 {
		return true, maxValue
	}

	return false, 0

}

func threeOfAKindCheak(allCard []PlayingСard) (bool, int8) {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	if ok, firstMaxValue := RepeatingCard(copyAllCard, 3); ok {

		copyAllCard = removeAllFindedCard(copyAllCard, firstMaxValue)

		if ok, secondMaxValue := RepeatingCard(copyAllCard, 3); ok {
			if firstMaxValue > secondMaxValue {
				return true, firstMaxValue
			} else {
				return true, secondMaxValue
			}
		} else {
			return false, firstMaxValue
		}
	} else {
		return false, 0
	}

}

func twoPairCheak(allCard []PlayingСard) (bool, int8, int8) {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	if ok, firstValue := RepeatingCard(copyAllCard, 2); ok {

		copyAllCard = removeAllFindedCard(copyAllCard, firstValue)

		if ok, secondValue := RepeatingCard(copyAllCard, 2); ok {

			copyAllCard = removeAllFindedCard(copyAllCard, secondValue)

			if ok, thirdValue := RepeatingCard(copyAllCard, 2); ok {

				allValue := []int8{firstValue, secondValue, thirdValue}

				allValue = sortIntInOrder(allValue)

				return true, allValue[2], allValue[1]
			}

			if firstValue > secondValue {
				return true, firstValue, secondValue
			} else {
				return true, secondValue, firstValue
			}

		} else {
			return false, 0, 0
		}

	} else {
		return false, 0, 0
	}

}

func sortIntInOrder(allValue []int8) []int8 {

	sort.SliceStable(allValue, func(i, j int) bool {
		return allValue[i] < allValue[j]
	})

	return allValue

}

func pairCheak(allCard []PlayingСard) (bool, int8) {

	copyAllCard := make([]PlayingСard, len(allCard))
	copy(copyAllCard, allCard)

	if ok, firstMaxValue := RepeatingCard(copyAllCard, 2); ok {
		return false, firstMaxValue
	} else {
		return false, 0
	}

}

func highCard(a, b int8) int8 {

	if a > b {
		return a
	} else {
		return b
	}

}

func winningPlayerInRoom(room Room) []Player {

	var lastPlayer Player
	var winningPlayer Player

	var allWinningPlayers []Player

	for _, p := range room {

		if p.Combination.Combination > lastPlayer.Combination.Combination {
			winningPlayer = p
		} else if p.Combination.Combination == lastPlayer.Combination.Combination {

			if p.Combination.FirstMaxValue > lastPlayer.Combination.FirstMaxValue {
				winningPlayer = p

			} else if p.Combination.FirstMaxValue == lastPlayer.Combination.FirstMaxValue {
				if p.Combination.SecondMaxValue > lastPlayer.Combination.SecondMaxValue {
					winningPlayer = p
				} else if p.Combination.SecondMaxValue == lastPlayer.Combination.SecondMaxValue {
					if p.Combination.Kicker > lastPlayer.Combination.Kicker {
						winningPlayer = p

					} else {
						if len(allWinningPlayers) == 0 {
							allWinningPlayers = append(allWinningPlayers, winningPlayer)
							allWinningPlayers = append(allWinningPlayers, p)
						} else {
							allWinningPlayers = append(allWinningPlayers, p)
						}
					}
				}
			}
		}

		lastPlayer = p
	}

	if len(allWinningPlayers) == 0 {
		allWinningPlayers = append(allWinningPlayers, winningPlayer)
	}

	return allWinningPlayers
}
