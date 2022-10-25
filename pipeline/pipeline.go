package pipeline

import (
	"encoding/csv"
	"log"
	"os"
	"sync"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/combinations"

	"gonum.org/v1/gonum/stat/combin"
)

func readCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

func writeCsv(filePath string, data [][]string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("error creating a file: %s", err)
	}
	writer := csv.NewWriter(file)
	for _, oneComb := range data {
		err = writer.Write(oneComb)
		if err != nil {
			log.Fatalf("error writing to a file: %s", err)
		}
	}
	writer.Flush()
	file.Close()
}

func removeDuplicates(cards []string) []string {
	list := []string{}
	allKeys := make(map[string]bool)
	for _, item := range cards {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func StrToCard(StrCard string) card.Card {
	strToBytes := []rune(StrCard)
	suit := string(strToBytes[0])
	face := string(strToBytes[1])

	return card.Card{Suit: suit, Face: face}
}

func getStructCards(cards []string) []card.Card {
	var records []card.Card
	for _, card := range cards {
		structCard := StrToCard(card)

		records = append(records, structCard)
	}
	return records
}

func checkCombinations(cards []card.Card) (combSlice []string) {
	var combName string
	switch true {
	case combinations.GetStraightFlush(cards):
		combName = " | Straight Flush"
	case combinations.GetFourOfAKind(cards):
		combName = " | Four Of A Kind"
	case combinations.GetFullHouse(cards):
		combName = " | Full House"
	case combinations.GetFlush(cards):
		combName = " | Flush"
	case combinations.GetStraight(cards):
		combName = " | Straight"
	case combinations.GetThreeOfAKind(cards):
		combName = " | Three Of A Kind"
	case combinations.GetTwoPairs(cards):
		combName = " | Two Pairs"
	case combinations.GetPair(cards):
		combName = " | Pair"
	default:
		combName = ""
	}
	if combName == "" {
		return []string{}
	}
	combSlice = structToSlice(cards)
	combSlice[4] += combName
	return combSlice
}

func structToSlice(comb []card.Card) []string {
	var res []string
	for _, cardItem := range comb {
		res = append(res, cardItem.Suit+cardItem.Face)
	}
	return res
}

func getCombName(cardsCombs [][]card.Card) [][]string {
	var result [][]string
	for _, comb := range cardsCombs {
		combName := checkCombinations(comb)
		if len(combName) == 0 {
			continue
		}
		result = append(result, combName)
	}
	return result
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

func Pipeline(num string, wg *sync.WaitGroup) {
	records := readCsv("dataset/dat" + num + ".csv")[0]
	records = removeDuplicates(records)
	cards := getStructCards(records)
	combin_result := getCombinations(cards, 5)
	combName := getCombName(combin_result)
	writeCsv("results/data"+num+".csv", combName)

	wg.Done()
}
