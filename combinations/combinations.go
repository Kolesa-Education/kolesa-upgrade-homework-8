package combinations

import (
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"sort"
)

type Combination struct {
	suit     int
	face1    int
	face2    int
	straight bool
	weight   int
}

var (
	Pair = Combination{
		face1:  2,
		weight: 2,
	}
	TwoPairs = Combination{
		face1:  2,
		face2:  2,
		weight: 3,
	}
	ThreeOfAKind = Combination{
		face1:  3,
		weight: 4,
	}
	Straight = Combination{
		straight: true,
		weight:   5,
	}
	Flush = Combination{
		suit:   5,
		weight: 6,
	}
	FullHouse = Combination{
		face1:  3,
		face2:  2,
		weight: 7,
	}
	FourOfAKind = Combination{
		face1:  4,
		weight: 8,
	}
	StraightFlush = Combination{
		suit:     5,
		straight: true,
		weight:   9,
	}
)

func GetFourOfAKind(cards []card.Card) (bool, []card.Card) {
	var (
		counter     int
		combination []card.Card
	)
	for i, curCard := range cards {
		combination = append(combination, curCard)
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				combination = append(combination, compCard)
				counter++
			}
			if counter == 3 {
				return true, combination
			}
		}
		counter = 0
		combination = []card.Card{}
	}
	return false, combination
}

func GetFullHouse(cards []card.Card) (bool, []card.Card) {
	var combination []card.Card
	check, threeOfAKind := GetThreeOfAKind(cards)
	if !check {
		return false, combination
	}
	combination = append(combination, threeOfAKind...)
	check, pair := GetPair(cards)
	if !check {
		combination = []card.Card{}
		return false, combination
	}
	combination = append(combination, pair...)
	return true, combination
}

func GetFlush(cards []card.Card) (bool, []card.Card) {
	var combination []card.Card
	firstCard := cards[0]
	combination = append(combination, firstCard)
	for _, compCard := range cards[+1:] {
		if firstCard.Suit != compCard.Suit {
			combination = []card.Card{}
			return false, combination
		}
		combination = append(combination, compCard)
	}
	return true, combination
}

func GetStraight(cards []card.Card) (bool, []card.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Weight < cards[j].Weight
	})
	var combination []card.Card
	firstCard := cards[0]
	lastCard := cards[len(cards)-1]
	if lastCard.Weight == 14 {
		check := checkAces(cards)
		if !check {
			return false, combination
		}
		lastCard.Weight = 1
		cardsWithoutAce := cards[:len(cards)-1]
		cardsAceFirst := make([]card.Card, 1)
		cardsAceFirst[0] = lastCard
		cardsAceFirst = append(cardsAceFirst, cardsWithoutAce...)
		fmt.Println(cardsAceFirst)
		check, aceFirst := GetStraight(cardsAceFirst)
		if check {
			return true, aceFirst
		}
	}
	combination = append(combination, firstCard)
	for i, compCard := range cards[1:] {
		if compCard.Weight != firstCard.Weight+(i+1) {
			combination = []card.Card{}
			return false, combination
		}
		combination = append(combination, compCard)
	}
	return true, combination
}

func checkAces(cards []card.Card) bool {
	cpCards := make([]card.Card, 0)
	copy(cpCards, cards)
	sort.Slice(cpCards, func(i, j int) bool {
		return cards[i].Weight > cards[j].Weight
	})
	firstAce := cards[0]
	for _, curCard := range cards[1:] {
		if firstAce.Weight == curCard.Weight {
			return false
		}
	}
	return true
}

func GetThreeOfAKind(cards []card.Card) (bool, []card.Card) {
	var (
		counter     int
		combination []card.Card
	)
	for i, curCard := range cards {
		combination = append(combination, curCard)
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				combination = append(combination, compCard)
				counter++
			}
			if counter == 2 {
				return true, combination
			}
		}
		counter = 0
		combination = []card.Card{}
	}
	return false, combination
}

func GetTwoPairs(cards []card.Card) (bool, []card.Card) {
	var (
		firstPair   bool
		combination []card.Card
	)
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				combination = append(combination, curCard)
				if !firstPair {
					firstPair = true
					combination = append(combination, compCard)
					continue
				}
				combination = append(combination, compCard)
				return true, combination
			}
		}
	}
	combination = []card.Card{}
	return false, combination
}

func GetPair(cards []card.Card) (bool, []card.Card) {
	var combination []card.Card
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				combination = append(combination, curCard, compCard)
				return true, combination
			}
		}
	}
	return false, combination
}
