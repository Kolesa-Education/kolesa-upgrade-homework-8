package utility

import (
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"strings"
)

func ConvertCards(content string) ([]card.Card, error) {
	arrStr := strings.Split(content, ",")
	var cards []card.Card
	var suit, face string

	for _, str := range arrStr {
		runes := []rune(str)

		if len(runes) == 2 || len(runes) != 3 {
			suit, face = string(runes[0]), string(runes[1])

		} else if len(runes) == 3 && string(runes[1:]) == "10" {
			suit, face = string(runes[0]), string(runes[1:])
		} else {
			return nil, fmt.Errorf("error in contetn of string")
		}

		cards = append(cards, card.Card{Face: face, Suit: suit})
	}
	return cards, nil
}
