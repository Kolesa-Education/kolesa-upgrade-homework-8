package service

import (
	"errors"
	"sort"
	"strconv"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"golang.org/x/exp/slices"
)

var errNotCombination = errors.New("Hmmmm. This is a not poker combination!")

// структура для
type PokerCombination struct {
	cards    []card.Card
	nameComb string
}

func (p *PokerCombination) getTrueCombination() error {
	if len(p.cards) != 5 {
		return errNotCombination
	}

	switch {
	case isStraightFlush(p.cards):
		p.nameComb = "Straight Flush"
	case isFourOfKind(p.cards):
		p.nameComb = "Four of a kind"
	case isFullHouse(p.cards):
		p.nameComb = "Full House"
	case isFlush(p.cards):
		p.nameComb = "Flush"
	case isStraight(p.cards):
		p.nameComb = "Straight"
	case isThreeOfKind(p.cards):
		p.nameComb = "Three of a Kind"
	case isTwoPairs(p.cards):
		p.nameComb = "Two Pairs"
	case isPair(p.cards):
		p.nameComb = "Pair"
	default:
		return errNotCombination
	}
	return nil
}

// функция которая считает количество номинала
func countFace(cards []card.Card) map[string]int {
	faceCount := map[string]int{}
	for _, elem := range cards {
		if count, ok := faceCount[elem.Face]; ok {
			faceCount[elem.Face] = count + 1
		} else {
			faceCount[elem.Face] = 1
		}
	}
	return faceCount
}

// функция которая считает количество мастей
func countSuit(cards []card.Card) map[string]int {
	suitCount := map[string]int{}
	for _, elem := range cards {
		if count, ok := suitCount[elem.Suit]; ok {
			suitCount[elem.Suit] = count + 1
		} else {
			suitCount[elem.Suit] = 1
		}
	}
	return suitCount
}

// функция для сортировки количества встречаемой масти или номинала,и возвращает массив состоящий из количеств мастей или номинала
func mapSortCounts(mapCount map[string]int) []int {
	result := make([]int, 0, len(mapCount))

	for i := range mapCount {
		result = append(result, mapCount[i])
	}

	sort.SliceStable(result, func(i, j int) bool { return result[i] > result[j] })
	return result

}

func isPair(card []card.Card) bool {
	sliceCount := mapSortCounts(countFace(card))
	if sliceCount[0] == 2 {
		return true
	}
	return false
}

func isTwoPairs(card []card.Card) bool {
	sliceCount := mapSortCounts(countFace(card))
	if sliceCount[0] == 2 && sliceCount[1] == 2 {
		return true
	}
	return false
}

func isThreeOfKind(card []card.Card) bool {
	sliceCount := mapSortCounts(countFace(card))
	if sliceCount[0] == 3 {
		return true
	}
	return false
}

func isStraight(card []card.Card) bool {
	var tempFace []int

	for _, value := range card {
		tempFace = append(tempFace, faceNum(value.Face))
	}

	sort.SliceStable(tempFace, func(i, j int) bool { return tempFace[i] < tempFace[j] })

	if slices.Contains(tempFace, 14) && slices.Contains(tempFace, 2) {
		if tempFace[0] == 2 && tempFace[1] == 3 && tempFace[2] == 4 && tempFace[3] == 5 && tempFace[4] == 14 {
			return true
		}
	}
	for i := range tempFace {
		if i != len(tempFace)-1 && tempFace[i]+1 != tempFace[i+1] {
			return false
		}
	}
	return true

}

func isFlush(card []card.Card) bool {
	sliceCount := mapSortCounts(countSuit(card))
	if sliceCount[0] == 5 {
		return true
	}
	return false
}

func isFullHouse(card []card.Card) bool {
	if isThreeOfKind(card) && isPair(card) {
		return true
	}
	return false
}

func isFourOfKind(card []card.Card) bool {
	sliceCount := mapSortCounts(countFace(card))
	if sliceCount[0] == 4 {
		return true
	}
	return false
}

func isStraightFlush(card []card.Card) bool {
	if isStraight(card) && isFlush(card) {
		return true
	}
	return false
}

func faceNum(s string) int {
	switch s {
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		n, _ := strconv.Atoi(s)
		return n
	}
}
