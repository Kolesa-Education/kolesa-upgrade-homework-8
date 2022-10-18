package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/combinations"
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

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func cardToStruct(cardStr string) (card.Card, error) {
	strToBytes := []rune(cardStr)
	suit := fmt.Sprintf("%q", strToBytes[0])
	face := fmt.Sprintf("%q", strToBytes[1])
	if suit == "" || face == "" {
		return card.Card{}, errors.New("cannot get cardStr suit and face")
	}
	return card.Card{Suit: suit, Face: face}, nil
}

func getStructCards(cards []string) []card.Card {
	var structCards []card.Card
	for _, c := range cards {
		cardStruct, err := cardToStruct(c)
		if err != nil {
			log.Fatalln("failed to convert string card representation to struct", err)
		}
		structCards = append(structCards, cardStruct)
	}
	return structCards
}

func main() {
	dataset, err := openDataset("./dataset/dattest.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	uniqCards := removeDuplicateStr(dataset)
	cards := getStructCards(uniqCards)
	fmt.Println(cards)
	check, pair := combinations.GetPair(cards)
	fmt.Println("Pair:", check, pair)
	check, twoPairs := combinations.GetTwoPairs(cards)
	fmt.Println("Two pairs:", check, twoPairs)
	check, threeOfAKind := combinations.GetThreeOfAKind(cards)
	fmt.Println("Three of a kind:", check, threeOfAKind)
	check, flush := combinations.GetFlush(cards)
	fmt.Println("Flush:", check, flush)
	check, fullHouse := combinations.GetFullHouse(cards)
	fmt.Println("Full House:", check, fullHouse)
}
