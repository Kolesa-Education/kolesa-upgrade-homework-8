package combinations

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"sort"
)

func IsPair(cards []card.Card) bool {
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				return true
			}
		}
	}
	return false
}

func IsTwoPairs(cards []card.Card) bool {
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

func IsThreeOfAKind(cards []card.Card) bool {
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

func IsStraight(cards []card.Card) bool {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Power < cards[j].Power
	})
	firstCard := cards[0]
	lastCard := cards[len(cards)-1]
	if lastCard.Power == 14 {
		check := IsAces(cards)
		if !check {
			return false
		}
		lastCard.Power = 1
		cardsWithoutAce := cards[:len(cards)-1]
		cardsAceFirst := make([]card.Card, 1)
		cardsAceFirst[0] = lastCard
		cardsAceFirst = append(cardsAceFirst, cardsWithoutAce...)
		aceFirst := IsStraight(cardsAceFirst)
		if aceFirst {
			return true
		}
	}
	for i, compCard := range cards[1:] {
		if compCard.Power != firstCard.Power+(i+1) {
			return false
		}
	}
	return true
}

func IsFlush(cards []card.Card) bool {
	firstCard := cards[0]
	for _, compCard := range cards[+1:] {
		if firstCard.Suit != compCard.Suit {
			return false
		}
	}
	return true
}

func IsFullHouse(cards []card.Card) bool {
	threeOfAKind := IsThreeOfAKind(cards)
	if !threeOfAKind {
		return false
	}
	pair := IsPair(cards)
	if !pair {
		return false
	}
	return true
}

func IsFourOfAKind(cards []card.Card) bool {
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

func IsStraightFlush(cards []card.Card) bool {
	straight := IsStraight(cards)
	if !straight {
		return false
	}
	flush := IsFlush(cards)
	if !flush {
		return false
	}
	return true
}

func IsAces(cards []card.Card) bool {
	cpCards := make([]card.Card, 0)
	copy(cpCards, cards)
	sort.Slice(cpCards, func(i, j int) bool {
		return cards[i].Power > cards[j].Power
	})
	firstAce := cards[0]
	for _, curCard := range cards[1:] {
		if firstAce.Power == curCard.Power {
			return false
		}
	}
	return true
}
