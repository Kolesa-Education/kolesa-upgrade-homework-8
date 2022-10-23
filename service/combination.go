package service

import combination "github.com/mxschmitt/golang-combinations"

const cardCombCount = 5

func getCombinationOfFiveCard(cards []string) [][]string {
	result := combination.Combinations(cards, cardCombCount)
	return result

}
