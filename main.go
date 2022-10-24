package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	combinations "github.com/mxschmitt/golang-combinations"
	"github.com/samber/lo"
)

func cardsToRepresentations(cards []card.Card) []string {
	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
		r, _ := c.ShortRepresentation()
		return r
	})
	return representations
}

func GetFiles() []fs.FileInfo {
	files, err := ioutil.ReadDir("dataset/")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func ReadCSV(fileName string) []string {
	f, err := os.Open("dataset/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dataCsv := csv.NewReader(f)
	data, err := dataCsv.Read()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ClearDuplicate(cardList []string) []string {
	tempMap := map[string]string{}
	for _, item := range cardList {
		tempMap[item] = "item"
	}
	cardList = cardList[:0]
	for key, _ := range tempMap {
		cardList = append(cardList, key)
	}

	return cardList
}

func GetCombinations(cardList []string) [][]string {
	result := combinations.Combinations(cardList, 5)
	return result

}

func getDenominations(cards []string) map[string]int {
	denom := []string{}
	for _, item := range cards {
		denom = append(denom, item[len(item)-1:])
	}
	result := lo.CountValues(denom)
	return result
}

func SortCardDenom(denom map[string]int) bool {
	reference := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	tempRefIndex := []int{}

	for index, refItem := range reference {
		for key, _ := range denom {
			if key == refItem {
				if len(tempRefIndex) == 0 {
					tempRefIndex = append(tempRefIndex, index)
				} else {
					if tempRefIndex[len(tempRefIndex)-1]+1 != index {
						return false
					} else {
						tempRefIndex = append(tempRefIndex, index)
					}
				}

			}
		}
	}

	if len(tempRefIndex) == 5 {
		return true
	} else {
		return false
	}
}

func getCombinationsType(cards []string) [][]string {
	denom := getDenominations(cards)
	var resCard [][]string
	switch len(denom) {
	case 4:
		resCard = append(resCard, cards)
		resCard = append(resCard, []string{"Pair"})
	case 3:
		for _, value := range denom {
			if value == 2 && len(resCard) == 0 {
				resCard = append(resCard, cards)
				resCard = append(resCard, []string{"Two Pairs"})
			}
			if value == 3 {
				resCard = append(resCard, cards)
				resCard = append(resCard, []string{"Three of a Kind"})
			}
		}

	case 2:
		for _, value := range denom {
			if value == 3 {
				resCard = append(resCard, cards)
				resCard = append(resCard, []string{"Full House"})
			} else if value == 4 {
				resCard = append(resCard, cards)
				resCard = append(resCard, []string{"Four of a kind"})
			}
		}

	case 5:
		tempMap := map[string]string{}
		for _, item := range cards {
			tempMap[item[:len(item)-1]] = "items"
		}

		if len(tempMap) == 1 && SortCardDenom(denom) {
			resCard = append(resCard, cards)
			resCard = append(resCard, []string{"Straight Flush"})
		}

		if len(tempMap) == 1 {
			resCard = append(resCard, cards)
			resCard = append(resCard, []string{"Flush"})
		}

		if SortCardDenom(denom) {
			resCard = append(resCard, cards)
			resCard = append(resCard, []string{"Straight"})
		}
	}

	return resCard
}

func SortByCombination(combination [][][]string) [][]string {
	combinationRef := []string{"Straight Flush", "Four of a kind", "Full House", "Flush", "Straight", "Three of a Kind", "Two Pairs", "Pair"}
	sortedMapCombinations := map[string][][]string{}
	sortedSliceCombi := [][]string{}
	for _, item := range combination {
		sortedMapCombinations[item[1][0]] = append(sortedMapCombinations[item[1][0]], item[0])
	}

	for _, combiName := range combinationRef {
		for key, value := range sortedMapCombinations {
			if key == combiName {
				for _, item := range value {
					fmt.Println(item)
					item = append(item, "|", key)
					tempString := strings.Join(item, " ")
					sortedSliceCombi = append(sortedSliceCombi, []string{tempString})
				}
			}
		}
	}
	return sortedSliceCombi
}

func WriteCsv(sortedCombination [][]string, index int) {

	csvFile, err := os.Create("result/data" + strconv.Itoa(index) + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	// for _, item := range sortedCombination {
	csvwriter.WriteAll(sortedCombination)
	// }

	csvwriter.Flush()
	csvFile.Close()
}

func CreateDir() {
	file, _ := os.Stat("result")

	if file != nil {
		os.RemoveAll("result")
	}

	if err := os.Mkdir("result", os.ModePerm); err != nil {
		log.Fatal(err)
	}

}

func BaseFunc(file fs.FileInfo, index int) {
	tempCombiSlice := [][][]string{}
	cardList := ReadCSV(file.Name())
	newCardList := ClearDuplicate(cardList)
	combi := GetCombinations(newCardList)
	for _, item := range combi {
		cards := getCombinationsType(item)
		if len(cards) != 0 {
			tempCombiSlice = append(tempCombiSlice, cards)
		}
	}

	finalList := SortByCombination(tempCombiSlice)

	WriteCsv(finalList, index)

}

func main() {
	var wg sync.WaitGroup

	CreateDir()
	filesList := GetFiles()
	wg.Add(len(filesList))
	for index, file := range filesList {
		go BaseFunc(file, index)

	}
	wg.Wait()

	// var seed int64 = 1665694295623135151
	// randomSource := rand.NewSource(seed)
	// random := rand.New(randomSource)
	// log.Printf("Initialized random with seed %d\n", seed)

	// fmt.Println("Starting to generate cards")
	// for i := 0; i < 100; i++ {
	// 	log.Printf("Iteration %d\n", i)
	// 	cardsInFile := random.Intn(7) + 10 // [10, 17]
	// 	cards := make([]card.Card, 0)

	// 	for j := 0; j < cardsInFile; j++ {
	// 		generatedCard, _ := card.Random(*random)
	// 		cards = append(cards, *generatedCard)
	// 	}
	// 	log.Printf("Generated cards %s\n", cards)
	// 	summary := cardsToRepresentations(cards)
	// 	file, err := os.Create(fmt.Sprintf("dataset/dat%d.csv", i))

	// 	if err != nil {
	// 		log.Fatalln("failed to open file", err)
	// 	}

	// 	writer := csv.NewWriter(file)
	// 	if err = writer.Write(summary); err != nil {
	// 		log.Fatalln("error writing to a file!")
	// 	}

	// 	writer.Flush()
	// 	_ = file.Close()
	// }
}

// pair 4

// Two Pairs 3
// Three of a Kind 3

// Four of a kind 2
// Full House 2

// Straight 5
// Flush 5
// Straight Flush	5
