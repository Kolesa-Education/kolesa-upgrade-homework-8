package hw

import (
	"errors"
	"sort"
	"strconv"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

var errNotCombination = errors.New("Not a Combination")

type CardCombination struct {
	Cards    []card.Card
	combName string
}

func (c CardCombination) GetCombinationName() string {
	return c.combName
}

func (c *CardCombination) DetectCombination() error {
	if len(c.Cards) != 5 {
		return errNotCombination
	}

	switch {
	case isStraightFlush(c.Cards):
		c.combName = "Straight Flush"
	case isFourKind(c.Cards):
		c.combName = "Four of a kind"
	case isFullHouse(c.Cards):
		c.combName = "Full House"
	case isFlush(c.Cards):
		c.combName = "Flush"
	case isStraight(c.Cards):
		c.combName = "Straight"
	case isThreeKind(c.Cards):
		c.combName = "Three of a Kind"
	case isTwoPairs(c.Cards):
		c.combName = "Two Pairs"
	case isPair(c.Cards):
		c.combName = "Pair"
	default:
		return errNotCombination
	}

	return nil
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
		if i == 0 && cardsFaces[i] == 0 && cardsFaces[i+1] == 11 {
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
