package pipeline

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/combinations"
	"gonum.org/v1/gonum/stat/combin"
	"log"
	"os"
	"regexp"
	"sync"
)

func openCSV(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file "+path, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.Read()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path, err)
	}
	return records, nil
}

func writeCSV(filePath string, data [][]string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvWriter := csv.NewWriter(f)
	for _, oneComb := range data {
		err = csvWriter.Write(oneComb)
		if err != nil {
			log.Fatalf("failed writing file: %s", err)
		}
	}
	csvWriter.Flush()
	f.Close()
}

func removeDuplicateStr(strSlice []string) []string {
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
	weight := getWeight(face)
	if suit == "" || face == "" {
		return card.Card{}, errors.New("cannot get cardStr data")
	}
	return card.Card{Suit: suit, Face: face, Weight: weight}, nil
}

func getWeight(face string) int {
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
			log.Fatalln("failed to convert string card representation to struct", err)
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

func combsToStrings(cardsCombs [][]card.Card) [][]string {
	var res [][]string
	for _, comb := range cardsCombs {
		combName := checkCombinations(comb)
		if combName == "" {
			continue
		}
		combLst := []string{combName}
		res = append(res, combLst)
	}
	return res
}

func checkCombinations(cards []card.Card) (combString string) {
	combString = structToString(cards)
	switch true {
	case combinations.GetStraightFlush(cards):
		combString += " | Straight Flush"
		return
	case combinations.GetFourOfAKind(cards):
		combString += " | Four Of A Kind"
		return
	case combinations.GetFullHouse(cards):
		combString += " | Full House"
		return
	case combinations.GetFlush(cards):
		combString += " | Flush"
		return
	case combinations.GetStraight(cards):
		combString += " | Straight"
		return
	case combinations.GetThreeOfAKind(cards):
		combString += " | Three Of A Kind"
		return
	case combinations.GetTwoPairs(cards):
		combString += " | Two Pairs"
		return
	case combinations.GetPair(cards):
		combString += " | Pair"
		return
	default:
		return ""
	}
}

func structToString(comb []card.Card) string {
	var res string
	for i, cardItem := range comb {
		res += fmt.Sprintf("%s%s", cardItem.Suit, cardItem.Face)
		if i != len(comb)-1 {
			res += ","
		}
	}
	return res
}

func removeQuotes(filename string) error {
	input2, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	output2 := regexp.MustCompile(`"`).ReplaceAll(input2, []byte(""))

	if err = os.WriteFile(filename, output2, 0666); err != nil {
		return err
	}
	return nil
}

func Pipeline(num string, group *sync.WaitGroup) {
	dataset, err := openCSV("./dataset/dat" + num + ".csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	uniqCards := removeDuplicateStr(dataset)
	cards := getStructCards(uniqCards)
	combStructLst := getCombinations(cards, 5)
	combStringLst := combsToStrings(combStructLst)
	filename := "./results/dat" + num + ".csv"
	writeCSV(filename, combStringLst)
	err = removeQuotes(filename)
	if err != nil {
		log.Fatalln("failed to remove quotes from file:", filename)
	}
	//fmt.Println(combStringLst)
	group.Done()
}
