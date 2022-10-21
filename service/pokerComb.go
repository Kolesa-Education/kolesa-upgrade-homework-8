package service

import (
	"errors"
	"sort"
	"strconv"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

const (
	CombinationSize          = 5
	CombinationPairName      = "Pair"
	CombinationTwoPairs      = "Two Pairs"
	CombinationThreeOfAKind  = "Three Of A Kind"
	CombinationStraight      = "Straight"
	CombinationFlush         = "Flush"
	CombinationFullHouse     = "Full House"
	CombinationFourOfAKind   = "Four Of A Kind"
	CombinationStraightFlush = "Straight Flush"
	NumericJack              = 11
	NumericQueen             = 12
	NumericKing              = 13
	NumericAce               = 14
)

type Combination struct {
	name  string
	cards []card.Card
}

func countFaces(cards []card.Card) map[string]int {
	faceCount := map[string]int{}
	for _, card := range cards {
		if count, ok := faceCount[card.Face]; ok {
			faceCount[card.Face] = count + 1
		} else {
			faceCount[card.Face] = 1
		}
	}
	return faceCount
}

func countSuits(cards []card.Card) map[string]int {
	suitsCount := map[string]int{}
	for _, card := range cards {
		if count, ok := suitsCount[card.Suit]; ok {
			suitsCount[card.Suit] = count + 1
		} else {
			suitsCount[card.Suit] = 1
		}
	}
	return suitsCount
}

func sortValues(mapComb map[string]int) []int {
	values := lo.Values[string, int](mapComb)
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	return values
}

func isCombOfPair(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 2
}

func isCombOfTwoPair(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 2 && sortValues(countFaces(cards))[1] == 2
}

func isCombOfThreeOfAKind(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 3
}

func numFace(cards string) int {
	switch cards {
	case card.FaceJack:
		return NumericJack
	case card.FaceQueen:
		return NumericQueen
	case card.FaceKing:
		return NumericKing
	case card.FaceAce:
		return NumericAce
	default:
		num, _ := strconv.Atoi(cards)
		return num
	}
}

func isCombOfStaright(cards []card.Card) bool {
	var sliceFace []int

	for _, value := range cards {
		sliceFace = append(sliceFace, numFace(value.Face))
	}

	sort.Slice(sliceFace, func(i, j int) bool { return sliceFace[i] > sliceFace[j] })

	if slices.Contains(sliceFace, NumericAce) && slices.Contains(sliceFace, 2) {
		return sliceFace[0] == NumericAce && sliceFace[1] == 5 && sliceFace[2] == 4 && sliceFace[3] == 3 && sliceFace[4] == 2
	} else {
		for i := 0; i < len(sliceFace); i++ {
			if i != len(sliceFace)-1 && sliceFace[i]-sliceFace[i+1] != 1 {
				return false
			}
		}
	}
	return true
}

func isCombOfFlush(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 5
}

func isCombFullHouse(cards []card.Card) bool {
	return isCombOfThreeOfAKind(cards) && isCombOfPair(cards)
}

func isCombFourOfAKind(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 4
}

func isCombOfStarightFlush(cards []card.Card) bool {
	return isCombOfStaright(cards) && isCombOfFlush(cards)
}

func (c *Combination) getTrueCombination() error {
	if len(c.cards) != 5 {
		return errors.New("Not correct")
	}
	switch {
	case isCombOfStarightFlush(c.cards):
		c.name = CombinationStraightFlush
	case isCombFourOfAKind(c.cards):
		c.name = CombinationFourOfAKind
	case isCombFullHouse(c.cards):
		c.name = CombinationFullHouse
	case isCombOfFlush(c.cards):
		c.name = CombinationFlush
	case isCombOfStaright(c.cards):
		c.name = CombinationStraight
	case isCombOfThreeOfAKind(c.cards):
		c.name = CombinationThreeOfAKind
	case isCombOfTwoPair(c.cards):
		c.name = CombinationTwoPairs
	case isCombOfPair(c.cards):
		c.name = CombinationPairName
	default:
		return errors.New("Not combination")
	}
	return nil
}
