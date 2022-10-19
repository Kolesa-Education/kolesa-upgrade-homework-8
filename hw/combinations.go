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

	fmt.Println(sortMapValues(countFaces(cards)))
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

func sortMapValues(m map[string]int) []int {
	values := make([]int, 0, len(m))

	for key := range m {
		values = append(values, m[key])
	}

	sort.SliceStable(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	return values
}

func isFaceComb(cards []card.Card, combNum1, combNum2 int) bool {
	sortedCardsFaces := sortMapValues(countFaces(cards))
	if sortedCardsFaces[0] == combNum1 && sortedCardsFaces[1] == combNum2 {
		return true
	}
	return false
}

func isPair(cards []card.Card) bool {
	return isFaceComb(cards, 2, 1)
}

func isTwoPairs(cards []card.Card) bool {
	return isFaceComb(cards, 2, 2)
}

func isThreeKind(cards []card.Card) bool {
	return isFaceComb(cards, 3, 1)
}

func isFullHouse(cards []card.Card) bool {
	return isFaceComb(cards, 3, 2)
}

func isFourKind(cards []card.Card) bool {
	return isFaceComb(cards, 4, 1)
}

func Straight(cards []card.Card) { // string {
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Face < cards[j].Face
	})
	fmt.Println(cards)
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
