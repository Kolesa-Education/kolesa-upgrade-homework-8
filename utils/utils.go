package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/comb"
	"gonum.org/v1/gonum/stat/combin"
)

func OpenCSV(path string) ([]string, error) {
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

func WriteIntoCSV(filePath string, data [][]string) {
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

func CardToStruct(cardStr string) (card.Card, error) {
	strToBytes := []rune(cardStr)
	suit := string(strToBytes[0])
	face := string(strToBytes[1])
	if face == "1" {
		face += "0"
	}
	weight := comb.GetWeight(face)
	if suit == "" || face == "" {
		return card.Card{}, errors.New("cannot get cardStr data")
	}
	return card.Card{Suit: suit, Face: face, Weight: weight}, nil
}

func FormatingCards(cardsFromFile []string) []card.Card {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range cardsFromFile {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	var structs []card.Card
	for _, c := range list {
		cardStruct, err := CardToStruct(c)
		if err != nil {
			log.Fatalln("failed to convert string card representation to struct", err)
		}
		structs = append(structs, cardStruct)
	}
	return structs
}

func GetCombinations(cards []card.Card, col int) [][]string {
	cs := combin.Combinations(len(cards), col)
	var superList [][]card.Card
	for _, c := range cs {
		var list []card.Card
		for i := 0; i < col; i++ {
			list = append(list, cards[c[i]])
		}
		superList = append(superList, list)
	}
	var res [][]string
	for _, comb := range superList {
		combName := CheckCombinations(comb)
		if len(combName) == 0 {
			continue
		}
		res = append(res, combName)
	}
	return res
}

func StructToString(comb []card.Card) string {
	var res string
	for i, cardItem := range comb {
		res += fmt.Sprintf("%s%s", cardItem.Suit, cardItem.Face)
		if i != len(comb)-1 {
			res += ","
		}
	}
	return res
}

func CheckCombinations(cards []card.Card) (combStrings []string) {
	combStrings = append(combStrings, StructToString(cards))
	switch true {
	case comb.GetStraightFlush(cards):
		combStrings = append(combStrings, " | Straight Flush")
		return
	case comb.GetFourOfAKind(cards):
		combStrings = append(combStrings, " | Four Of A Kind")
		return
	case comb.GetFullHouse(cards):
		combStrings = append(combStrings, " | Full House")
		return
	case comb.GetFlush(cards):
		combStrings = append(combStrings, " | Flush")
		return
	case comb.GetStraight(cards):
		combStrings = append(combStrings, " | Straight")
		return
	case comb.GetThreeOfAKind(cards):
		combStrings = append(combStrings, " | Three Of A Kind")
		return
	case comb.GetTwoPairs(cards):
		combStrings = append(combStrings, " | Two Pairs")
		return
	case comb.GetPair(cards):
		combStrings = append(combStrings, " | Pair")
		return
	default:
		return []string{}
	}
}
