package service

import combo "github.com/mxschmitt/golang-combinations"

func CombinationCards(cards []string) [][]string {
	return combo.Combinations(cards, CombinationSize)
}
