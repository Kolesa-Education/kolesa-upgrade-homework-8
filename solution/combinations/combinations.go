package combinations

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/solution/utility"
)

/* Рекурсивная функция который обработывает комбинаций
inputArr ---> Входной слайс с картами
outputArr ---> Временный массив с размером 5 карт для хранения текущей комбинации
start & end ---> Начальный и конечный индексы inputArr
index ---> Текущий индекс в outputArr
r ---> Размер комбинации
*/

func Combinations(inputArr []card.Card, outputArr [5]card.Card,
	start, end, index, r int, outputFile string) {
	if index == r {
		Process(outputArr, outputFile) // Обработка одной комбинации
		return
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
		outputArr[index] = inputArr[i]
		Combinations(inputArr, outputArr, i+1, end, index+1, r, outputFile)
	}
}

// Проверяет карты на покерую комбинацию с наиболее сильной комбинаций

func Process(cards [5]card.Card, filename string) {
	var handType = ""
	if IsStraightFlush(cards) {
		handType = "straight flush"
	} else if IsFourKind(cards) {
		handType = "four kind"
	} else if IsFullHouse(cards) {
		handType = "full house"
	} else if IsFlush(cards) {
		handType = "flush"
	} else if IsStraight(cards) {
		handType = "straight"
	} else if IsThreeKind(cards) {
		handType = "three kind"
	} else if IsTwoPairs(cards) {
		handType = "two pairs"
	} else if IsPair(cards) {
		handType = "pair"
	} else {
		return
	}
	if handType != "" {
		utility.ToFile(cards, handType, filename) // Записывает в файл
	}
}
