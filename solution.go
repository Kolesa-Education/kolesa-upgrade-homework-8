package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()
	return file
}

func getDataFromCSV(fileName string) []string {
	file := readFile(fileName)
	csvReader := csv.NewReader(file)
	data, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}
	return data

}

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
		keys[i] = string(k[valueIndex])
		i++
	}
	fmt.Println(dataMap)
	return keys
}

func checkQuantativeCombinations(valuesSlice []string, dataMap map[string]int) string {
	quantativeCombinations := map[string]string{
		"22": "2 pairs",
		"23": "full house",
		"32": "full house",
		"2":  "pair",
		"3":  "three of a kind",
		"4":  "four of a kind",
	}
	res := ""
	combination := ""
	for i := 0; i < len(valuesSlice); i++ {
		//fmt.Println(dataMap[valuesSlice[i]])
		switch dataMap[valuesSlice[i]] {
		case 2:
			res += "2"
		case 3:
			res += "3"
		case 4:
			res += "4"
		}
	}
	fmt.Println(res)
	res, ok := quantativeCombinations[combination]

	if ok {
		return res
	}
	return "0"

}

func getCardValuesSliceAndSuitsMap(dataSlice []string) ([]string, map[string]bool) {
	var valuesSlice []string = make([]string, len(dataSlice))
	suitsMap := make(map[string]bool)
	for i := 0; i < len(dataSlice); i++ {
		suitsMap[string(dataSlice[i][0:3])] = true
		valuesSlice[i] = string(dataSlice[i][3])
	}
	return valuesSlice, suitsMap
}

func checkStraightOrFlush(valuesSlice []string, suitsMap map[string]bool) string {
	combination := ""
	cardsOrder := "A2345678910JQKA"
	if strings.Contains(cardsOrder, strings.Join(valuesSlice, "")) {
		combination += "straight"
	}

	if len(suitsMap) == 1 {
		combination += "flush"
	}
	return combination
}

func main() {
	fileName := "dataset/dat0.csv"
	dataSlice := getDataFromCSV(fileName)
	dataMap := getCardMapFromSlice(dataSlice)
	values := getUniqueValuesFromDataMap(dataMap, 3)
	uniqueValuesMap := getCardMapFromSlice(values)
	uniqueValues := getUniqueValuesFromDataMap(uniqueValuesMap, 0)

	_, suitsMap := getCardValuesSliceAndSuitsMap(dataSlice)
	sort.Strings(uniqueValues)
	fmt.Println(uniqueValues)
	//fmt.Println(dataMap)
	//fmt.Println(dataMap)
	isFlushOrStraight := checkStraightOrFlush(uniqueValues, suitsMap)
	isQuantative := checkQuantativeCombinations(uniqueValues, uniqueValuesMap)
	fmt.Print(isFlushOrStraight + " " + isQuantative)

}
