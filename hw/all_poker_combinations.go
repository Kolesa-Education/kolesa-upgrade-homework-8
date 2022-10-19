package hw

import (
	"fmt"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

type Result struct {
	fileName string
	result   string
}

func (r Result) GetFileName() string {
	return r.fileName
}

func (r Result) GetResult() string {
	return r.result
}

func GetAllPokerCombinations(fileName string, ch chan Result) {
	cardsStrSlice, err := getCardsStrSlice(fileName)
	if err != nil {
	}

	allCombinations := getAllCardCombinations(cardsStrSlice)

	var cardCombinationsResult string
	for _, comb := range allCombinations {
		cards := getCardsFromStr(comb)
		cardComb := CardCombination{
			Cards: cards,
		}
		if err := cardComb.DetectCombination(); err != nil {
			continue
		}

		cardCombinationsResult += fmt.Sprintf("%s| %s\n", cardsToStr(cards), cardComb.GetCombinationName())
	}

	result := Result{
		fileName: fileName,
		result:   cardCombinationsResult,
	}
	ch <- result
}

func cardsToStr(cards []card.Card) string {
	var result string

	for i := range cards {
		result += fmt.Sprintf("%s%s ", cards[i].Suit, cards[i].Face)
	}

	return result
}
