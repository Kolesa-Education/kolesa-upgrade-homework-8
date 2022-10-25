package main

import (
	"errors"
	"sort"
	"strings"
)

func getCardMapFromSlice(dataSlice []string) map[string]int {
	dataMap := make(map[string]int)
	for i := 0; i < len(dataSlice); i++ {
		dataMap[dataSlice[i]]++
	}
	return dataMap
}

func getUniqueValuesFromDataMap(dataMap map[string]int, valueIndex int) []string {
	keys := make([]string, len(dataMap))
	i := 0
	for k := range dataMap {
		keys[i] = k
		if valueIndex != -1 {
			var val string = string(k[valueIndex])
			if val == "1" {
				val = "10"
			}
			keys[i] = string(val)
		}

		i++
	}
	return keys
}

func checkQuantativeCombinations(valuesSlice []string, dataMap map[string]int) (string, error) {
	quantativeCombinations := map[string]string{
		"22": "Two pairs",
		"23": "Full House",
		"32": "Full House",
		"2":  "Pair",
		"3":  "Three Of A Kind",
		"4":  "Four Of A Kind",
	}
	res := ""
	for i := 0; i < len(valuesSlice); i++ {
		switch dataMap[valuesSlice[i]] {
		case 2:
			res += "2"
		case 3:
			res += "3"
		case 4:
			res += "4"
		}
	}
	combination, ok := quantativeCombinations[res]

	if ok {
		return combination, nil
	}
	return "", errors.New("Empty combination")

}

func getSuitsMap(dataSlice []string) map[string]bool {
	suitsMap := make(map[string]bool)
	for i := 0; i < len(dataSlice); i++ {
		suitsMap[string(dataSlice[i][0:3])] = true
	}
	return suitsMap
}

func cardStringToInt(card string) string {
	highSuits := map[string]string{
		"10": "A",
		"J":  "B",
		"Q":  "C",
		"K":  "D",
		"A":  "E",
	}
	res, ok := highSuits[card]
	if ok {
		return res
	}
	return card
}

func transformCards(cards []string) []string {
	for i := range cards {
		cards[i] = cardStringToInt(cards[i])
	}
	return cards
}

func checkStraightOrFlush(uniqueValues []string, suitsMap map[string]bool) string {
	if len(uniqueValues) != 5 {
		return ""
	}
	//sort uniqueValues by function
	combination := ""
	cardsOrder := "E123456789ABCDE"
	uniqueValues = transformCards(uniqueValues)
	sort.Slice(uniqueValues, func(i, j int) bool {
		return uniqueValues[i] < uniqueValues[j]
	})

	combinationString := strings.Join(uniqueValues, "")
	var alternativeCombinationString string = "PLOT_TWIST"

	if strings.Contains(combinationString, "E") {
		alternativeCombinationString = "E" + strings.Join(uniqueValues[:4], "")
	}

	var isSubstring bool = strings.Contains(cardsOrder, combinationString)
	isAlternativeSubstring := strings.Contains(cardsOrder, alternativeCombinationString)
	var isStraigth bool = isSubstring || isAlternativeSubstring

	if isStraigth {
		combination += "Straight"

	}
	if len(suitsMap) == 1 {
		combination += "Flush"
	}

	return combination
}

func findCombination(dataSlice []string) string {
	dataMap := getCardMapFromSlice(dataSlice)

	cardValues := getUniqueValuesFromDataMap(dataMap, 3)

	cardValuesMap := getCardMapFromSlice(cardValues)

	uniqueValues := getUniqueValuesFromDataMap(cardValuesMap, 0)

	suitsMap := getSuitsMap(dataSlice)

	isQuantative, err := checkQuantativeCombinations(uniqueValues, cardValuesMap)
	if err != nil {
		isQuantative = ""
	}

	isFlushOrStraight := checkStraightOrFlush(uniqueValues, suitsMap)

	resultantCombination := isQuantative
	if isFlushOrStraight != "" {
		resultantCombination = isFlushOrStraight
	}
	if isFlushOrStraight == "" && isQuantative == "" {
		return ""
	}

	return strings.Join(dataSlice, ",") + " | " + resultantCombination
}
