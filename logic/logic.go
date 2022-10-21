package logic

import (
	"encoding/csv"
	"fmt"
	"sync"

	"sort"

	//"errors"
	//"fmt"
	//"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	//"github.com/Kolesa-Education/kolesa-upgrade-homework-8/combinations"
	//"gonum.org/v1/gonum/stat/combin"
	"log"
	"os"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/combinations"
	"gonum.org/v1/gonum/stat/combin"
	//"sync"
)

func ReadLineCSV(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("OPENING ERROR "+path, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.Read()
	if err != nil {
		log.Fatal("CONVERATION ERROR "+path, err)
	}

	return records, nil
}

func WriteAllLinesCSV(path string, data []string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("CREATION ERROR %s", err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	err = csvWriter.Write(data)
	if err != nil {
		log.Fatalf("WRITING ERROR  %s", err)
	}

	csvWriter.Flush()

}

func RemoveDuplicates(strSlice []string) []string {
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

func ToPower(face string) int {
	switch face {
	case "2":
		return card.Power2
	case "3":
		return card.Power3
	case "4":
		return card.Power4
	case "5":
		return card.Power5
	case "6":
		return card.Power6
	case "7":
		return card.Power7
	case "8":
		return card.Power8
	case "9":
		return card.Power9
	case "10":
		return card.Power10
	case "J":
		return card.PowerJack
	case "Q":
		return card.PowerQueen
	case "K":
		return card.PowerKing
	case "A":
		return card.PowerAce
	default:
		return 0
	}

}

func GetCombinations(cards []card.Card, cardNumber int) [][]card.Card {

	idxcombinations := combin.Combinations(len(cards), cardNumber)

	var resList [][]card.Card

	for _, c := range idxcombinations {
		var list []card.Card
		for i := 0; i < cardNumber; i++ {
			list = append(list, cards[c[i]])

		}
		resList = append(resList, list)

	}
	return resList

}

func ToStringArray(cards [][]card.Card) (res []string) {

	sort.Slice(cards, func(i, j int) bool {
		return checkCombination(cards[i]) < checkCombination(cards[j])

	})

	for i := 0; i < len(cards); i++ {
		strCons := ""

		for j := 0; j < len(cards[i]); j++ {
			strCons += ToString(cards[i][j])
		}

		combName := getCombinationName(checkCombination(cards[i]))
		if combName != "" {
			strCons += " | " + combName + "\n"
			res = append(res, strCons)

		}

	}
	return res

}

func getCombinationName(combPower int) string {

	switch combPower {
	case 0:
		return "Pair"
	case 1:
		return "TwoPairs"
	case 2:
		return "ThreeOfaKind"
	case 3:
		return "Straight"
	case 4:
		return "Flush"
	case 5:
		return "FullHouse"
	case 6:
		return "FourOfaKind"
	case 7:
		return "StraightFlush"
	default:
		return ""
	}

}

func checkCombination(comb []card.Card) int {

	switch {

	case combinations.IsStraightFlush(comb):
		return 7
	case combinations.IsFourOfaKind(comb):
		return 6
	case combinations.IsFullHouse(comb):
		return 5
	case combinations.IsFlush(comb):
		return 4
	case combinations.IsStraight(comb):
		return 3
	case combinations.IsThreeOfaKind(comb):
		return 2
	case combinations.IsTwoPairs(comb):
		return 1
	case combinations.IsPair(comb):
		return 0
	default:
		return -1
	}

}

func ToCard(cardStr string) card.Card {
	cardRunes := []rune(cardStr)
	var suit string = ""
	var face string = ""
	for i, s := range cardRunes {
		if i == 0 {
			suit = string(s)
		} else {
			face += string(s)
		}
	}

	return card.Card{Suit: suit, Face: face, Power: ToPower(face)}

}

func ToString(card card.Card) string {
	return fmt.Sprintf("%s%s", card.Suit, card.Face)
}

func ExecuteMain(num string, group *sync.WaitGroup) {

	data, _ := ReadLineCSV("dataset/dat" + num + ".csv")
	withoutdupl := RemoveDuplicates(data)
	mas := []card.Card{}

	for _, str := range withoutdupl {
		mas = append(mas, ToCard(str))
	}

	allcomb := GetCombinations(mas, 5)

	stringList := ToStringArray(allcomb)

	WriteAllLinesCSV("results/dat"+num+".csv", stringList)
	group.Done()
}
