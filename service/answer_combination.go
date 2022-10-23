package service

import (
	"fmt"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

type Answer struct {
	filePath string
	answer   string
}

// func (a Answer) GetFileName() string {
// 	return a.filePath
// }

func (a Answer) GetAnswer() string {
	return a.answer
}

// функция которая соединяет в себе все функции для отправки ответа в main
func GetAnswerPokerCombination(filepath string, ch chan Answer) {
	cardsStrSlice, err := parseCardsSlice(filepath)
	if err != nil {

	}

	allComb := getCombinationOfFiveCard(cardsStrSlice)

	var cardCombinationAnswer string

	for _, comb := range allComb {
		cards := getCardFromStr(comb)
		cardComb := PokerCombination{
			cards: cards,
		}
		if err := cardComb.getTrueCombination(); err != nil {
			continue
		}
		cardCombinationAnswer += fmt.Sprintf("%s|%s\n", getCardsToStr(cards), cardComb.nameComb)
	}

	answer := Answer{
		filePath: filepath,
		answer:   cardCombinationAnswer,
	}
	ch <- answer
}

func getCardsToStr(cards []card.Card) string {
	var result string

	for i := range cards {
		result += fmt.Sprintf("%s%s", cards[i].Suit, cards[i].Face)
	}

	return result
}
