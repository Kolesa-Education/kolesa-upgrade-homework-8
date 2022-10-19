package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/combinations"
	"gonum.org/v1/gonum/stat/combin"
	"log"
	"os"
)

//func cardsToRepresentations(cards []card.Card) []string {
//	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
//		r, _ := c.ShortRepresentation()
//		return r
//	})
//	return representations
//}

func openDataset(path string) ([]string, error) {
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

//func Contains(a []string, x string) bool {
//	for _, n := range a {
//		if x == n {
//			return true
//		}
//	}
//	return false
//}

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

func main() {
	dataset, err := openDataset("./dataset/dat0.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	uniqCards := removeDuplicateStr(dataset)
	cards := getStructCards(uniqCards)
	combsList := getCombinations(cards, 5)
	combLst := combsToStrings(combsList)
	fmt.Println(combLst)
}

func combsToStrings(cardsCombs [][]card.Card) []string {
	var stringLst []string
	for _, comb := range cardsCombs {
		combName := checkCombinations(comb)
		stringLst = append(stringLst, combName)
	}
	return stringLst
}

func checkCombinations(cards []card.Card) string {
	var combString string
	check, straightFlush := combinations.GetStraightFlush(cards)
	if check {
		combString = structToString(cards)
		combString += fmt.Sprintf(" | Straight Flush")
		return combString
	}
	check, fourOfAKind := combinations.GetFourOfAKind(cards)
	if check {
		combString = structToString(straightFlush)
		combString += fmt.Sprintf(" | Straight Flush")
		return combString
	}
	fmt.Println("Four Of A Kind:", check, fourOfAKind)
	check, fullHouse := combinations.GetFullHouse(cards)
	fmt.Println("Full House:", check, fullHouse)
	check, flush := combinations.GetFlush(cards)
	fmt.Println("Flush:", check, flush)
	check, straight := combinations.GetStraight(cards)
	fmt.Println("Straight:", check, straight)
	check, threeOfAKind := combinations.GetThreeOfAKind(cards)
	fmt.Println("Three Of A Kind:", check, threeOfAKind)
	check, twoPairs := combinations.GetTwoPairs(cards)
	fmt.Println("Two pairs:", check, twoPairs)
	check, pair := combinations.GetPair(cards)
	fmt.Println("Pair:", check, pair)
}

func structToString(comb []card.Card) string {
	var res string
	for _, card := range comb {
		res += fmt.Sprintf("%s%s", card.Suit, card.Face)
	}
	return res
}
