package combinations

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"sort"
)

var cardOrderDict = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

func IsStraightFlush(cards [5]card.Card) bool {
	if IsFlush(cards) && IsStraight(cards) {
		return true
	}
	return false
}

func IsFourKind(cards [5]card.Card) bool {
	faceCounter := make(map[string]int)
	for _, card := range cards {
		faceCounter[card.Face]++
	}

	var values []int
	for _, value := range faceCounter {
		values = append(values, value)
	}
	sort.Ints(values)

	if values[len(values)-1] == 4 {
		return true
	}
	return false
}

func IsFullHouse(cards [5]card.Card) bool {
	faceCounter := make(map[string]int)
	for _, card := range cards {
		faceCounter[card.Face]++
	}

	var values []int
	for _, value := range faceCounter {
		values = append(values, value)
	}
	sort.Ints(values)

	if len(values) == 2 && values[0] == 2 && values[1] == 3 {
		return true
	}
	return false
}

func IsFlush(cards [5]card.Card) bool {
	m := make(map[string]bool)

	for i := 0; i < len(cards); i++ {
		suit := cards[i].Suit
		m[suit] = true
	}
	if len(m) == 1 {
		return true
	}
	return false
}

func IsStraight(cards [5]card.Card) bool {
	faceCounter := make(map[string]bool)
	var ranks []int
	for _, card := range cards {
		if faceCounter[card.Face] {
			return false
		}
		faceCounter[card.Face] = true
		ranks = append(ranks, cardOrderDict[card.Face])
	}

	sort.Ints(ranks)

	if len(faceCounter) == 5 && ranks[len(ranks)-1]-ranks[0] == 4 {
		return true
	}

	for _, r := range "A2345" {
		if !faceCounter[string(r)] {
			return false
		}
	}

	return true
}

func IsThreeKind(cards [5]card.Card) bool {
	faceCounter := make(map[string]int)
	for _, card := range cards {
		faceCounter[card.Face]++
	}

	var values []int
	for _, value := range faceCounter {
		values = append(values, value)
	}
	sort.Ints(values)

	if values[len(values)-1] == 3 {
		return true
	}
	return false
}

func IsTwoPairs(cards [5]card.Card) bool {
	faceCounter := make(map[string]int)
	for _, card := range cards {
		faceCounter[card.Face]++
	}

	var values []int
	for _, value := range faceCounter {
		values = append(values, value)
	}
	sort.Ints(values)

	if len(values) == 3 && values[0] == 1 && values[1] == 2 && values[2] == 2 {
		return true
	}
	return false
}

func IsPair(cards [5]card.Card) bool {
	faceCounter := make(map[string]int)

	for _, card := range cards {
		faceCounter[card.Face]++
	}

	for _, counts := range faceCounter {
		if counts == 2 {
			return true
		}
	}
	return false
}
