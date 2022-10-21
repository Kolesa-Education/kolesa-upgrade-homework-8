package service

import (
	"fmt"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

type Answer struct {
	filePath string
	answer   string
}

func (a Answer) GetAnswer() string {
	return a.answer
}

func GetAnswerPokerCombination(filepath string, ch chan Answer) {
	cardsStrSlice, err := ReadFile(filepath)
	if err != nil {

	}

	allComb := CombinationCards(cardsStrSlice)

	var cardCombinationAnswer string

	for _, comb := range allComb {
		cards := ConvertToCard(comb)
		cardComb := Combination{
			cards: cards,
		}
		if err := cardComb.getTrueCombination(); err != nil {
			continue
		}
		cardCombinationAnswer += fmt.Sprintf("%s|%s\n", getCardsToStr(cards), cardComb.name)
		//fmt.Println(cardCombinationAnswer)
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
