package service

import (
	"io/ioutil"
	"strings"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

// функция собирает другие функции внутри и отдает массив карт
func finalParseCards(filepath string) ([]card.Card, error) {
	cardStr, err := parseCardsSlice(filepath)
	if err != nil {
		return nil, err
	}
	return getCardFromStr(cardStr), nil
}

// функция читает из файла, удаляет дубликаты и записывает в срез
func parseCardsSlice(filepath string) ([]string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	tempSlice := string(data)
	cardsSlice := deleteDuplicate(strings.Split(tempSlice[:len(tempSlice)-1], ","))

	return cardsSlice, nil

}

// функция отделяет масть и номинал по полям структуры
func getCardFromStr(strCard []string) []card.Card {
	var results []card.Card
	var card card.Card
	for _, w := range strCard {
		for i, c := range w {
			if checkSuit(string(c)) {
				card.Suit = string(c)
			} else {
				card.Face = w[i:]
			}
		}
		results = append(results, card)
	}
	return results

}

// функция для удаления дубликатов из слайса
func deleteDuplicate(cards []string) []string {
	tempMap := make(map[string]bool)
	result := []string{}
	for _, item := range cards {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = true
			result = append(result, item)
		}
	}
	return result
}

// функция проверяет масть карты
func checkSuit(suit string) bool {
	switch suit {
	case card.SuitClubsUnicode, card.SuitDiamondsUnicode, card.SuitHeartsUnicode, card.SuitSpadesUnicode:
		return true
	default:
		return false
	}
}
