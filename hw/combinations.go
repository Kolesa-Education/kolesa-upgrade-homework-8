package hw

import (
	"errors"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

var errNotCombination = errors.New("Not a Combination")

type CardComb struct {
	cards []card.Card
}

func isSuit(suit string) bool {
	switch suit {
	case card.SuitClubsUnicode, card.SuitDiamondsUnicode, card.SuitHeartsUnicode, card.SuitSpadesUnicode:
		return true
	default:
		return false
	}
}

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

func countFaces(cards []card.Card) map[string]int {
	result := make(map[string]int)
	for _, card := range cards {
		result[card.Face] += 1
	}
	return result
}

func countSuits(cards []card.Card) map[string]int {
	result := make(map[string]int)

	for _, card := range cards {
		result[card.Suit] += 1
	}
	return result
}

func sortMapValues(m map[string]int) []int {
	values := make([]int, 0, len(m))

	for key := range m {
		values = append(values, m[key])
	}

	sort.SliceStable(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	return values
}

func isFaceComb(cards []card.Card, combNum1, combNum2 int) bool {
	sortedCardsFaces := sortMapValues(countFaces(cards))
	if sortedCardsFaces[0] == combNum1 && sortedCardsFaces[1] == combNum2 {
		return true
	}
	return false
}

func isPair(cards []card.Card) bool {
	return isFaceComb(cards, 2, 1)
}

func isTwoPairs(cards []card.Card) bool {
	return isFaceComb(cards, 2, 2)
}

func isThreeKind(cards []card.Card) bool {
	return isFaceComb(cards, 3, 1)
}

func isFullHouse(cards []card.Card) bool {
	return isFaceComb(cards, 3, 2)
}

func isFourKind(cards []card.Card) bool {
	return isFaceComb(cards, 4, 1)
}

func isStraight(cards []card.Card) bool {
	var cardsFaces []int

	for _, card := range cards {
		cardsFaces = append(cardsFaces, getFaceNum(card.Face))
	}

	sort.SliceStable(cardsFaces, func(i, j int) bool {
		return cardsFaces[i] < cardsFaces[j]
	})

	for i := range cardsFaces {
		if i == 0 && cardsFaces[i] == 0 {
			continue
		}

		if i != len(cardsFaces)-1 && cardsFaces[i]+1 != cardsFaces[i+1] {
			return false
		}
	}

	return true
}

func getFaceNum(s string) int {
	switch s {
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 0
	default:
		n, _ := strconv.Atoi(s)
		return n
	}
}

func isFlush(cards []card.Card) bool {
	suitCount := countSuits(cards)

	if len(suitCount) == 1 {
		return true
	}

	return false
}

func isStraightFlush(cards []card.Card) bool {
	if isFlush(cards) && isStraight(cards) {
		return true
	}
	return false
}

func DetectCombination(cards []card.Card) (string, error) {
	if len(cards) != 5 {
		return "", errNotCombination
	}

	switch {
	case isStraightFlush(cards):
		return "Straight Flush", nil
	case isFourKind(cards):
		return "Four of a kind", nil
	case isFullHouse(cards):
		return "Full House", nil
	case isFlush(cards):
		return "Flush", nil
	case isStraight(cards):
		return "Straight", nil
	case isThreeKind(cards):
		return "Three of a Kind", nil
	case isTwoPairs(cards):
		return "Two Pairs", nil
	case isPair(cards):
		return "Pair", nil
	default:
		return "", errNotCombination
	}
}
