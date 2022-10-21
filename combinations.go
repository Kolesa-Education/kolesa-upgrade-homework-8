package main

import "github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"

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
