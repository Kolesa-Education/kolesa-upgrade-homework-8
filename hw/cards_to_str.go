package hw

import (
	"fmt"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

func CardsToStr(cards []card.Card) string {
	var result string

	for i := range cards {
		result += fmt.Sprintf("%s%s ", cards[i].Suit, cards[i].Face)
	}

	return result
}
