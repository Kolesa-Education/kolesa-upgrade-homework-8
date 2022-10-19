package combinations

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"sort"
)

func GetStraightFlush(cards []card.Card) bool {
	straight := GetStraight(cards)
	if !straight {
		return false
	}
	flush := GetFlush(cards)
	if !flush {
		return false
	}
	return true
}

func GetFourOfAKind(cards []card.Card) bool {
	var counter int
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				counter++
			}
			if counter == 3 {
				return true
			}
		}
		counter = 0
	}
	return false
}

func GetFullHouse(cards []card.Card) bool {
	threeOfAKind := GetThreeOfAKind(cards)
	if !threeOfAKind {
		return false
	}
	pair := GetPair(cards)
	if !pair {
		return false
	}
	return true
}

func GetFlush(cards []card.Card) bool {
	firstCard := cards[0]
	for _, compCard := range cards[+1:] {
		if firstCard.Suit != compCard.Suit {
			return false
		}
	}
	return true
}

func GetStraight(cards []card.Card) bool {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Weight < cards[j].Weight
	})
	firstCard := cards[0]
	lastCard := cards[len(cards)-1]
	if lastCard.Weight == 14 {
		check := checkAces(cards)
		if !check {
			return false
		}
		lastCard.Weight = 1
		cardsWithoutAce := cards[:len(cards)-1]
		cardsAceFirst := make([]card.Card, 1)
		cardsAceFirst[0] = lastCard
		cardsAceFirst = append(cardsAceFirst, cardsWithoutAce...)
		aceFirst := GetStraight(cardsAceFirst)
		if aceFirst {
			return true
		}
	}
	for i, compCard := range cards[1:] {
		if compCard.Weight != firstCard.Weight+(i+1) {
			return false
		}
	}
	return true
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

func GetThreeOfAKind(cards []card.Card) bool {
	var counter int
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				counter++
			}
			if counter == 2 {
				return true
			}
		}
		counter = 0
	}
	return false
}

func GetTwoPairs(cards []card.Card) bool {
	var firstPair bool
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				if !firstPair {
					firstPair = true
					continue
				}
				return true
			}
		}
	}
	return false
}

func GetPair(cards []card.Card) bool {
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				return true
			}
		}
	}
	return false
}
