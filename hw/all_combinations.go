package hw

import combinations "github.com/mxschmitt/golang-combinations"

const cardsNumberInComb = 5

func getAllCardCombinations(cardsStr []string) [][]string {
	result := combinations.Combinations(cardsStr, cardsNumberInComb)
	return result
}
