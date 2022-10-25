package procedures

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Card struct {
	Suit  string
	Face  string
	Power int
}

var powerCard = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

func findStraight(cards []Card) []Card {
	fmt.Println(cards)
	res := make([]Card, 0)
	deckIndex := len(cards) - 1

	lastCard := cards[deckIndex]
	previousCard := cards[deckIndex-1]
	//[{hearts 3 3} {spades 6 6} {diamonds 7 7} {hearts 7 7} {diamonds 9 9} {clubs 9 9} {diamonds J 11} {clubs J 11} {hearts Q 12} {spades Q 12} {clubs K 13} {hearts A 14}]
	for deckIndex >= 1 {
		//fmt.Println(res)
		//fmt.Println(deckIndex)
		//fmt.Println("-------------")

		if lastCard.Power == previousCard.Power {
			deckIndex--
			previousCard = cards[deckIndex-1]
			continue
		}
		if lastCard.Power-1 == previousCard.Power {
			res = append(res, lastCard)
			deckIndex--
			lastCard = previousCard
			previousCard = cards[deckIndex-1]
		} else {
			res = make([]Card, 0)
			lastCard = previousCard
			deckIndex--
		}
		if len(res) == 5 {
			break
		}
	}
	return res
}

//func findPair(cards []Card) []Card {
//
//}

var stringifySuitCard = map[string]string{
	"\u2666": "diamonds",
	"\u2663": "clubs",
	"\u2665": "hearts",
	"\u2660": "spades",
}

func Run(pathToFile string) {
	f := DelDuplicate(ReadCsv(pathToFile))
	res := convToCardStruct(f)
	SortCards(res)
	cs := findStraight(res)
	fmt.Print("Result-----------")
	fmt.Println(SortCards(cs))
}

func SortCards(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Power < cards[j].Power {
			return true
		}
		return false
	})

	return cards
}

func convToCardStruct(cards []string) []Card {
	res := make([]Card, 0)
	for _, c := range cards {
		arr := strings.SplitAfterN(c, "", 2)
		res = append(res, Card{
			Suit:  stringifySuitCard[arr[0]],
			Face:  arr[1],
			Power: powerCard[arr[1]],
		})
	}

	return res
}

func DelDuplicate(cardSet []string) []string {
	res := make([]string, 0)
	for _, card := range cardSet {
		if !contains(res, card) {
			res = append(res, card)
		}
	}

	return res
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func GetFilesNames(pathToDir string) []string {
	files, err := ioutil.ReadDir(pathToDir)
	if err != nil {
		log.Fatal(err)
	}

	filesNames := make([]string, 0)
	for _, file := range files {
		filesNames = append(filesNames, fmt.Sprintf("%s/%s", pathToDir, file.Name()))
	}

	return filesNames
}

func ReadCsv(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	data, _ := reader.Read()

	return data
}
