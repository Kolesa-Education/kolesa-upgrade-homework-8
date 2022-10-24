package pipeline

import (
	"encoding/csv"
	"errors"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/combinations"
	"gonum.org/v1/gonum/stat/combin"
	"log"
	"os"
	"sync"
)

func readCsvFile(filePath string) ([]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.Read()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records, nil
}

func writeCsvFile(filePath string, data [][]string) {
	csvFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	for _, oneComb := range data {
		err = csvWriter.Write(oneComb)
		if err != nil {
			log.Fatalf("failed writing file: %s", err)
		}
	}
	csvWriter.Flush()
	csvFile.Close()
}

func removeDuplicates(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func cardToStruct(cardStr string) (card.Card, error) {
	strToBytes := []rune(cardStr)
	suit := string(strToBytes[0])
	face := string(strToBytes[1])
	if face == "1" {
		face += "0"
	}
	power := getPower(face)
	if suit == "" || face == "" {
		return card.Card{}, errors.New("cannot get cardStr data")
	}
	return card.Card{Suit: suit, Face: face, Power: power}, nil
}

func getPower(face string) int {
	switch face {
	case card.Face2:
		return 2
	case card.Face3:
		return 3
	case card.Face4:
		return 4
	case card.Face5:
		return 5
	case card.Face6:
		return 6
	case card.Face7:
		return 7
	case card.Face8:
		return 8
	case card.Face9:
		return 9
	case card.Face10:
		return 10
	case card.FaceJack:
		return 11
	case card.FaceQueen:
		return 12
	case card.FaceKing:
		return 13
	case card.FaceAce:
		return 14
	default:
		return 0
	}
}

func getStructCards(cards []string) []card.Card {
	var structs []card.Card
	for _, c := range cards {
		cardStruct, err := cardToStruct(c)
		if err != nil {
			log.Fatal(err)
		}
		structs = append(structs, cardStruct)
	}
	return structs
}

func getCombinations(cards []card.Card, col int) [][]card.Card {
	cs := combin.Combinations(len(cards), col)
	var superList [][]card.Card
	for _, c := range cs {
		var list []card.Card
		for i := 0; i < col; i++ {
			list = append(list, cards[c[i]])
		}
		superList = append(superList, list)
	}
	return superList
}

func combinationsToStrings(cardsCombs [][]card.Card) [][]string {
	var res [][]string
	for _, comb := range cardsCombs {
		combinationName := checkCombinations(comb)
		if len(combinationName) == 0 {
			continue
		}
		res = append(res, combinationName)
	}
	return res
}

func checkCombinations(cards []card.Card) (combSlice []string) {
	var combinationName string
	switch true {
	case combinations.IsStraightFlush(cards):
		combinationName = ", | Straight Flush"
	case combinations.IsFourOfAKind(cards):
		combinationName = ", | Four Of A Kind"
	case combinations.IsFullHouse(cards):
		combinationName = ", | Full House"
	case combinations.IsFlush(cards):
		combinationName = ", | Flush"
	case combinations.IsStraight(cards):
		combinationName = ", | Straight"
	case combinations.IsThreeOfAKind(cards):
		combinationName = ", | Three Of A Kind"
	case combinations.IsTwoPairs(cards):
		combinationName = ", | Two Pairs"
	case combinations.IsPair(cards):
		combinationName = ", | Pair"
	default:
		combinationName = ""
	}
	if combinationName == "" {
		return []string{}
	}
	combSlice = structToSlice(cards)
	combSlice[4] += combinationName
	return combSlice
}

func structToSlice(comb []card.Card) []string {
	var res []string
	for _, cardItem := range comb {
		res = append(res, cardItem.Suit+cardItem.Face)
	}
	return res
}

func Pipeline(num string, wg *sync.WaitGroup) {
	dataset, err := readCsvFile("./dataset/dat" + num + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	datasetCards := removeDuplicates(dataset)
	uniqueCards := getStructCards(datasetCards)
	combinationsStructList := getCombinations(uniqueCards, 5)
	combinationsStringList := combinationsToStrings(combinationsStructList)
	filename := "./results/data" + num + ".csv"
	writeCsvFile(filename, combinationsStringList)
	wg.Done()
}
