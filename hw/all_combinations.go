package hw

import combinations "github.com/mxschmitt/golang-combinations"

const cardsNumberInComb = 5

func GetAllCardCombinations(cardsStr []string) [][]string {
	result := combinations.Combinations(cardsStr, cardsNumberInComb)
	return result
}
