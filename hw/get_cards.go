package hw

import (
	"io/ioutil"
	"strings"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

func GetCards(fileName string) ([]card.Card, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	s := string(content)
	strCards := removeDuplicateStr(strings.Split(s[:len(s)-1], ","))

	var cards []card.Card
	var card card.Card
	for _, c := range strCards {
		for i, v := range c {
			if isSuit(string(v)) {
				card.Suit = string(v)
			} else {
				card.Face = c[i:]
				break
			}
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func isSuit(suit string) bool {
	switch suit {
	case card.SuitClubsUnicode, card.SuitDiamondsUnicode, card.SuitHeartsUnicode, card.SuitSpadesUnicode:
		return true
	default:
		return false
	}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
