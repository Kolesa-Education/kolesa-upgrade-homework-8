package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	comb "github.com/Kolesa-Education/kolesa-upgrade-homework-8/combinations"
	"github.com/samber/lo"
	"log"
	"os"
	"sync"
)

func cardsToRepresentations(cards []card.Card) []string {
	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
		r, _ := c.ShortRepresentation()
		return r
	})
	return representations
}

func readCSV(filename string) ([]string, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	reader := csv.NewReader(file)
	records, err := reader.Read()
	if err != nil {
		return nil, errors.New("failed to parse file as CSV")
	}
	return records, nil
}

func deleteDuplicates(dataList []string) []string {
	var newList []string
	for _, v1 := range dataList {
		counter := 0
		for _, v2 := range newList {
			if v1 == v2 {
				counter++
			}
		}
		if counter == 0 {
			newList = append(newList, v1)
		}
	}
	return newList
}

func dataToCards(data []string) []card.Card {
	var cards []card.Card
	for i := 0; i < len(data); i++ {
		cards[i].Suit = string(data[i][0])
		cards[i].Face = string(data[i][1])
	}
	return cards
}

func makeCombinations(cards []card.Card) [][]card.Card {
	count := 0
	var combinations [][]card.Card
	for i := 0; i < len(cards); i++ {
		combinations[count][0] = cards[i]
		for j := 0; i < len(cards); j++ {
			if j != i {
				combinations[count][1] = cards[j]
			}
			for k := 0; i < len(cards); k++ {
				if k != j && k != i {
					combinations[count][2] = cards[k]
				}
				for n := 0; i < len(cards); n++ {
					if n != j && n != i && n != k {
						combinations[count][3] = cards[n]
						count++
						combinations[count] = combinations[count-1]
					}
					for m := 0; i < len(cards); m++ {
						if m != j && m != i && m != k && m != n {
							combinations[count][4] = cards[m]
							count++
							combinations[count] = combinations[count-1]
						}
					}
				}
			}
		}
	}
	return combinations
}

func writeCsv(n int, records [][]string) error {
	file, err := os.Create(fmt.Sprintf("results/dat%d.csv", n))
	defer file.Close()
	if err != nil {
		return errors.New("failed to open file")
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return errors.New("failed writing to a file")
		}
	}
	return nil
}

func run(n int) {
	record, err := readCSV(fmt.Sprintf("dataset/dat%d.csv", n))
	record = deleteDuplicates(record)
	if err != nil {
		log.Fatalf(err.Error())
	}
	cards := dataToCards(record)
	combinations := makeCombinations(cards)
	var res [][]string
	for i := len(combinations); i < len(combinations); i++ {
		switch comb.GetCombination(cards) {
		case comb.Pair:
			res[i] = append(res[i], cardsToRepresentations(combinations[i])[0]+"|"+comb.Pair)
		case comb.TwoPair:
			res[i] = append(res[i], cardsToRepresentations(combinations[i])[0]+"|"+comb.TwoPair)
		case comb.ThreeOfKind:
			res[i] = append(res[i], cardsToRepresentations(combinations[i])[0]+"|"+comb.ThreeOfKind)
		case comb.Straight:
			res[i] = append(res[i], cardsToRepresentations(combinations[i])[0]+"|"+comb.Straight)
		case comb.Flush:
			res[i] = append(res[i], cardsToRepresentations(combinations[i])[0]+"|"+comb.Flush)
		case comb.FullHouse:
			res[i] = append(res[i], cardsToRepresentations(combinations[i])[0]+"|"+comb.FullHouse)
		case comb.FourOfKind:
			res[i] = append(res[i], cardsToRepresentations(combinations[i])[0]+"|"+comb.FourOfKind)
		case comb.StraightFlush:
			res[i] = append(res[i], cardsToRepresentations(combinations[i])[0]+"|"+comb.StraightFlush)
		default:
			continue
		}
	}
	err = writeCsv(n, res)
	if err != nil {
		log.Fatal("failed write to file")
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(100)

	defer waitGroup.Wait()

	for i := 0; i < 100; i++ {
		go func() {
			defer waitGroup.Done()
			go run(i)
		}()
	}
}

//func main() {
//	var seed int64 = 1665694295623135151
//	randomSource := rand.NewSource(seed)
//	random := rand.New(randomSource)
//	log.Printf("Initialized random with seed %d\n", seed)
//
//	fmt.Println("Starting to generate cards...")
//	for i := 0; i < 100; i++ {
//		log.Printf("Iteration %d\n", i)
//		cardsInFile := random.Intn(7) + 10 // [10, 17]
//		cards := make([]card.Card, 0)
//
//		for j := 0; j < cardsInFile; j++ {
//			generatedCard, _ := card.Random(*random)
//			cards = append(cards, *generatedCard)
//		}
//		log.Printf("Generated cards %s\n", cards)
//		summary := cardsToRepresentations(cards)
//		file, err := os.Create(fmt.Sprintf("dataset/dat%d.csv", i))
//
//		if err != nil {
//			log.Fatalln("failed to open file", err)
//		}
//
//		writer := csv.NewWriter(file)
//		if err = writer.Write(summary); err != nil {
//			log.Fatalln("error writing to a file!")
//		}
//
//		writer.Flush()
//		_ = file.Close()
//	}
//}
