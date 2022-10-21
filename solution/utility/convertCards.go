package utility

import (
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"strings"
)

func ConvertToCards(content string) ([]card.Card, error) {
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
			return nil, fmt.Errorf("error in content of string")
		}

		cards = append(cards, card.Card{Face: face, Suit: suit})
	}
	return cards, nil
}

func UniqueCards(cards []card.Card) []card.Card {
	check := make(map[string]bool)
	var newCards []card.Card

	for _, card := range cards {
		str := card.Suit + card.Face
		if check[str] {
			continue
		} else {
			newCards = append(newCards, card)
			check[str] = true
		}
	}

	return newCards
}
