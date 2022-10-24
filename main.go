package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/samber/lo"
)

func cardsToRepresentations(cards []card.Card) []string {
	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
		r, _ := c.ShortRepresentation()
		return r
	})
	return representations
}

func createCardSet(data [][]string) []card.Card {
	var cardSet []card.Card
	for _, line := range data {
		var rec card.Card
		for _, i := range line {
			for j := 0; j < len(i); j++ {
				// improvisation, it does not work
				strToBytes := []rune(i)
				rec.Face = string(strToBytes[0])
				rec.Suit = string(strToBytes[1])
			}
			fmt.Print(rec)
		}
		cardSet = append(cardSet, rec)
	}

	return cardSet
}

func readFile(path string) []card.Card {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var cardSet = createCardSet(data)

	return cardSet
}

func readAllFiles() [][]card.Card {
	var setOfCardSet [][]card.Card
	for i := 0; i < 100; i++ {
		file := "dataset/dat" + strconv.Itoa(i) + ".csv"

		_ = append(setOfCardSet, readFile(file))
	}

	return setOfCardSet
}

func checkIfPair(cardSet []card.Card) bool {
	// O(n^2) ужасно, но не настолько
	for i := 0; i < len(cardSet); i++ {
		for j := i + 1; j < len(cardSet); j++ {
			if cardSet[i].Face == cardSet[j].Face {
				return true
			}
		}
	}

	return false
}

func checkIfDoublePair(cardSet []card.Card) bool {
	first := false
	second := false

	for i := 0; i < len(cardSet); i++ {
		for j := i + 1; j < len(cardSet); j++ {
			if cardSet[i].Face == cardSet[j].Face {
				if first {
					second = true
				} else {
					first = true
				}
			}
		}
	}

	return first && second
}

func checkIfThreeOfAKind(cardSet []card.Card) bool {
	// O(n^3) О БОЖЕ
	for i := 0; i < len(cardSet); i++ {
		for j := i + 1; j < len(cardSet); j++ {
			for k := j + 1; j < len(cardSet); k++ {
				if cardSet[i].Face == cardSet[j].Face && cardSet[i].Face == cardSet[k].Face {
					return true
				}
			}
		}
	}

	return false
}

func faceToRank(Face string) int {
	if Face == "2" {
		return 1
	}

	if Face == "3" {
		return 2
	}

	if Face == "4" {
		return 3
	}

	if Face == "5" {
		return 4
	}

	if Face == "6" {
		return 5
	}

	if Face == "7" {
		return 6
	}

	if Face == "8" {
		return 7
	}

	if Face == "9" {
		return 8
	}

	if Face == "10" {
		return 9
	}

	if Face == "J" {
		return 10
	}

	if Face == "Q" {
		return 10
	}

	if Face == "K" {
		return 10
	}

	if Face == "A" {
		return 11
	}

	return 0
}

func checkIfStraight(cardSet []card.Card) bool {
	sort.Slice(cardSet, func(i, j int) bool {
		return faceToRank(cardSet[i].Face) < faceToRank(cardSet[j].Face)
	})

	init := faceToRank(cardSet[0].Face)

	// ne handlil A 2 3 4 5
	for i := 1; i < len(cardSet); i++ {
		rank := faceToRank(cardSet[i].Face) - 1
		if rank-1 != init {
			return false
		}
	}

	return true
}

func checkIfFlush(cardSet []card.Card) bool {
	for i := 1; i < len(cardSet); i++ {
		if cardSet[i].Suit != cardSet[i-1].Suit {
			return false
		}
	}

	return true
}

func checkIfFullHouse(cardSet []card.Card) bool {
	for i := 6; i < len(cardSet); i++ {
		temp := cardSet[i-6 : i]
		return checkIfPair(temp) && checkIfThreeOfAKind(temp)
	}

	return false
}

func checkIfFourOfAKind(cardSet []card.Card) bool {
	// O(n^4) no comments
	for i := 0; i < len(cardSet); i++ {
		for j := i + 1; j < len(cardSet); j++ {
			for k := j + 1; k < len(cardSet); k++ {
				for l := k + 1; l < len(cardSet); l++ {
					if cardSet[i] == cardSet[j] && cardSet[i] == cardSet[k] && cardSet[i] == cardSet[l] {
						return true
					}
				}
			}
		}
	}

	return false
}

func checkIfFiveOfAKind(cardSet []card.Card) bool {
	// O(n^4) no comments
	for i := 0; i < len(cardSet); i++ {
		for j := i + 1; j < len(cardSet); j++ {
			for k := j + 1; k < len(cardSet); k++ {
				for l := k + 1; l < len(cardSet); l++ {
					for m := l + 1; m < len(cardSet); m++ {
						if cardSet[i] == cardSet[j] && cardSet[i] == cardSet[k] && cardSet[i] == cardSet[l] && cardSet[i] == cardSet[m] {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

func checkIfStraightFlush(cardSet []card.Card) bool {
	for i := 6; i < len(cardSet); i++ {
		temp := cardSet[i-6 : i]
		return checkIfFiveOfAKind(temp) && checkIfFlush(temp)
	}

	return false
}

// func checkCombinations(setOfCardSet [][]card.Card) {
// 	for i := 0; i < len(setOfCardSet); i++ {
// 		if(checkIfStraightFlush(setOfCardSet[i])){
// 		}
// 	}
// }

func main() {
	// var seed int64 = 1665694295623135151
	// randomSource := rand.NewSource(seed)
	// random := rand.New(randomSource)
	// log.Printf("Initialized random with seed %d\n", seed)

	// fmt.Println("Starting to generate cards...")
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

	readAllFiles()
}
