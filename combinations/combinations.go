package combinations

import (
	"strings"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

// For ranking by face
const Order = "23456789TJQKA"

func removeDuplicates(hand []string) []string {
	result := []string{}
	cards := make(map[string]bool)
	for _, item := range hand {
		if _, ok := cards[item]; !ok {
			cards[item] = true
			result = append(result, item)
		}
	}
	return result
}

func Len(cards []card.Card) int {
	return 5
}

func isFlush(cards []card.Card) bool {
	firstSuite := cards[0]
	for _, card := range cards {
		if card.Suit != firstSuite.Suit {
			return false
		}
	}
	return true
}

func isStraight(cards []card.Card) bool {
	var values []byte
	for _, card := range cards {
		values = append(values, card[1])
	}
	if strings.Contains(Order, string(values)) {
		return true
	}
	return false
}

func isStraightFlush(cards []card.Card) bool {
	if isStraight(cards) {
		if isFlush(cards) {
			return true
		}
		return isStraight(cards)
	}
	if isFlush(cards) {
		return isFlush(cards)
	}
	return false
}

func isFullHouse(cards []card.Card) bool {
	if isThreeOfAKind(cards) {
		if isPair(cards) {
			return true
		}
		return isThreeOfAKind(cards)
	}
	if isPair(cards) {
		return isPair(cards)
	}
	return false
}

func isFourOfAKind(cards []card.Card) bool {
	var counter int
	for i, card := range cards {
		for _, icard := range cards[i+1:] {
			if card == icard {
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

func isPair(cards []card.Card) bool {
	for i, card := range cards {
		for _, icard := range cards[i+1:] {
			if card == icard {
				return true
			}
		}
	}
	return false
}

func isTwoPairs(cards []card.Card) bool {
	var onePair bool
	for i, card := range cards {
		for _, icard := range cards[i+1:] {
			if card == icard {
				if !onePair {
					onePair = true
					continue
				}
				return true
			}
		}
	}
	return false
}

func isThreeOfAKind(cards []card.Card) bool {
	var counter int
	for i, card := range cards {
		for _, icard := range cards[i+1:] {
			if card == icard {
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
