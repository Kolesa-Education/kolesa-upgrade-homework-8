package service

import (
	"io/ioutil"
	"strings"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

func ReadFile(fileName string) ([]string, error) {
	fileRead, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	tempSlice := string(fileRead)
	deduplicate := unique(strings.Split(tempSlice[:len(tempSlice)-1], ","))
	return deduplicate, nil

}

func unique(cards []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range cards {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func ConvertToCard(cards []string) []card.Card {
	var results []card.Card
	var tempCard card.Card
	for _, w := range cards {
		for i, c := range w {
			if checkSuit(string(c)) {
				tempCard.Suit = string(c)
			} else {
				tempCard.Face = w[i:]
			}
		}
		results = append(results, tempCard)
	}
	return results
}

func checkSuit(s string) bool {
	switch s {
	case card.SuitClubsUnicode, card.SuitDiamondsUnicode, card.SuitHeartsUnicode, card.SuitSpadesUnicode:
		return true
	}
	return false
}
