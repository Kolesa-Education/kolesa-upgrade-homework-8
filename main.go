package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/stat/combin"
)

func cardsToRepresentations(cards []card.Card) []string {
	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
		r, _ := c.ShortRepresentation()
		return r
	})
	return representations
}

func readCsvFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.Read()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func writeCsvFile(records []string, i int) {
	f, err := os.Create(fmt.Sprintf("results/data%d.csv", i))
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write(records); err != nil {
		log.Fatalln("error writing record to file", err)
	}
}

func removeDuplicates(records []string) []string {
	allKeys := make(map[string]bool)
	var distinctCards []string
	for _, i := range records {
		if _, value := allKeys[i]; !value {
			allKeys[i] = true
			distinctCards = append(distinctCards, i)
		}
	}
	return distinctCards
}

func getCardCombinations(n int) {
	records := readCsvFile(fmt.Sprintf("dataset/dat%d.csv", n))
	distinctrecords := removeDuplicates(records)
	var splitcards []string

	cs := combin.Combinations(len(distinctrecords), 5)
	for _, c := range cs {
		fmt.Printf("%s,%s,%s,%s,%s\n", distinctrecords[c[0]], distinctrecords[c[1]], distinctrecords[c[2]], distinctrecords[c[3]], distinctrecords[c[4]])
	}

	size := 1
	var j int
	for i := 0; i < len(distinctrecords); i += size {
		j += size
		if j > len(distinctrecords) {
			j = len(distinctrecords)
		}
		distinctrecordsString := strings.Join(distinctrecords[i:j], "")
		splitcard := strings.SplitN(distinctrecordsString, "", 2)
		switch splitcard[0] {
		case "\u2666":
			splitcard[0] = "diamonds"
		case "\u2663":
			splitcard[0] = "clubs"
		case "\u2665":
			splitcard[0] = "hearts"
		case "\u2660":
			splitcard[0] = "spades"
		}
		switch splitcard[1] {
		case "J":
			splitcard[1] = "11"
		case "Q":
			splitcard[1] = "12"
		case "K":
			splitcard[1] = "13"
		case "A":
			splitcard[1] = "14"
		}
		splitcardString := strings.Join(splitcard, "")
		splitcards = append(splitcards, splitcardString)
	}

	fmt.Println(distinctrecords)
	fmt.Println(splitcards)
	writeCsvFile(distinctrecords, n)
}

func main() {
	var seed int64 = 1665694295623135151
	randomSource := rand.NewSource(seed)
	random := rand.New(randomSource)
	log.Printf("Initialized random with seed %d\n", seed)

	fmt.Println("Starting to generate cards...")
	for i := 0; i < 100; i++ {
		log.Printf("Iteration %d\n", i)
		cardsInFile := random.Intn(7) + 10 // [10, 17]
		cards := make([]card.Card, 0)

		for j := 0; j < cardsInFile; j++ {
			generatedCard, _ := card.Random(*random)
			cards = append(cards, *generatedCard)
		}
		log.Printf("Generated cards %s\n", cards)
		summary := cardsToRepresentations(cards)
		file, err := os.Create(fmt.Sprintf("dataset/dat%d.csv", i))

		if err != nil {
			log.Fatalln("failed to open file", err)
		}

		writer := csv.NewWriter(file)
		if err = writer.Write(summary); err != nil {
			log.Fatalln("error writing to a file!")
		}

		writer.Flush()
		_ = file.Close()
	}

	for i := 0; i < 100; i++ {
		getCardCombinations(i)
	}
}
