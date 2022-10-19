package hw

import (
	"fmt"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

func GetAllPokerCombinations(dataFileName, resultFileName string) {
	cardsStrSlice, err := getCardsStrSlice(dataFileName)
	if err != nil {
	}

	allCombinations := getAllCardCombinations(cardsStrSlice)

	for _, comb := range allCombinations {
		cards := getCardsFromStr(comb)
		cardComb := CardCombination{
			Cards: cards,
		}
		if err := cardComb.DetectCombination(); err != nil {
			continue
		}

		result := fmt.Sprintf("%s| %s", cardsToStr(cards), cardComb.GetCombinationName())
		fmt.Println(result)
	}
}

func cardsToStr(cards []card.Card) string {
	var result string

	for i := range cards {
		result += fmt.Sprintf("%s%s ", cards[i].Suit, cards[i].Face)
	}

	return result
}
