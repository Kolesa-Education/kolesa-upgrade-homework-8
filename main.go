package main

import (
	"encoding/csv"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	_ "github.com/natemcintosh/gocombinatorics"
	"github.com/samber/lo"
	"log"
	"math/rand"
	"os"
	_ "reflect"
	"sort"
	"time"
)

func cardsToRepresentations(cards []card.Card) []string {
	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
		r, _ := c.ShortRepresentation()
		return r
	})
	return representations
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
		formatCards(cards)

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
}

func formatCards(cards []card.Card) {
	//лютый говнокод
	go formatPair(cards)
	go formatTwoPairs(cards)

}

func formatTwoPairs(card []card.Card) {
	var m = make(map[int]string)
	var y = make(map[string]int)
	var cmbntn = []string{}

	for i := 0; i < len(card); i++ {
		m[i] = card[i].Suit
	}
	sortedArr := []string{}
	for i := 0; i < len(m); i++ {
		sortedArr = append(sortedArr, m[i])
	}

	sort.Strings(sortedArr)

	for i := 1; i < len(sortedArr); i++ {

		if sortedArr[i-1] == sortedArr[i] {
			cmbntn = append(cmbntn, sortedArr[i-1])
			cmbntn = append(cmbntn, sortedArr[i])
			break
		}
	}
	for i := 1; i < len(m); i++ {
		if m[i] != m[i-1] && cmbntn[0] != m[i-1] {
			cmbntn = append(cmbntn, m[i-1])
			if len(cmbntn) == 5 {
				break
			}
		}
	}
	for i := 0; i < len(cmbntn); i++ {
		y[cmbntn[i]] = i
	}
	if len(y) == 3 {
		fmt.Println(cmbntn)
	}

}
func formatPair(card []card.Card) {

	var m = make(map[int]string)
	var y = make(map[string]int)
	var cmbntn = []string{}

	for i := 0; i < len(card); i++ {
		m[i] = card[i].Suit
		y[card[i].Suit] = i
	}
	sortedArr := []string{}
	for i := 0; i < len(m); i++ {
		sortedArr = append(sortedArr, m[i])
	}

	sort.Strings(sortedArr)

	for i := 1; i < len(sortedArr); i++ {

		if sortedArr[i-1] == sortedArr[i] {
			cmbntn = append(cmbntn, sortedArr[i-1])
			cmbntn = append(cmbntn, sortedArr[i])
			break
		}
	}

	for key := range y {
		if cmbntn[0] != key {
			cmbntn = append(cmbntn, key)
		}
	}

	if len(cmbntn) == 5 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cmbntn), func(i, j int) { cmbntn[i], cmbntn[j] = cmbntn[j], cmbntn[i] })
		fmt.Println(cmbntn)
	}

}
