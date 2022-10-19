package hw

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

type CardComb struct {
	cards []card.Card
}

func isSuit(suit string) bool {
	switch suit {
	case card.SuitClubsUnicode, card.SuitDiamondsUnicode, card.SuitHeartsUnicode, card.SuitSpadesUnicode:
		return true
	default:
		return false
	}
}

func GetCards(fileName string) ([]card.Card, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	s := string(content)
	strCards := removeDuplicateStr(strings.Split(s[:len(s)-1], ","))

	var cards []card.Card
	var card card.Card
	for _, c := range strCards {
		for i, v := range c {
			if isSuit(string(v)) {
				card.Suit = string(v)
			} else {
				card.Face = c[i:]
				break
			}
		}
		cards = append(cards, card)
	}

	fmt.Println(sortMapByValue(countFaces(cards)))
	return cards, nil
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func countFaces(cards []card.Card) map[string]int {
	result := make(map[string]int)
	for _, card := range cards {
		result[card.Face] += 1
	}
	return result
}

func countSuits(cards []card.Card) map[string]int {
	result := make(map[string]int)

	for _, card := range cards {
		result[card.Suit] += 1
	}
	return result
}

func sortMapByValue(m map[string]int) map[string]int {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return m
}

func Straight(cards []card.Card) { // string {
	sortCards(cards)
	fmt.Println(cards)
}

func sortCards(table []card.Card) {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if getNum(table[j].Face) > getNum(table[j+1].Face) {
				temp := table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
}

func getNum(s string) int {
	switch s {
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 0
	default:
		n, _ := strconv.Atoi(s)
		return n
	}
}
