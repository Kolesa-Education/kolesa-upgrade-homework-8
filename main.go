package main

import (
	"encoding/csv"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/pockerCombination"
	"github.com/samber/lo"
	"log"
	"math/rand"
	"os"
)

func generateAll(cards []card.Card, index int, res []card.Card, id int) {
	if index == len(cards) {
		return
	}

	if len(res) == 5 {
		combination := pockerCombination.Combination{
			res,
		}
		if title, ok := combination.GetTitle(); ok {
			fmt.Println("%s %s\n", res, title)
			summary := cardsToRepresentations(res)
			size := len(summary)
			summary[size-1] = summary[size-1] + " | " + title
			file, err := os.Create(fmt.Sprintf("dataset/dat%d.csv", id))
			if err != nil {
				log.Fatalln("failed to open file", err)
			}

			writer := csv.NewWriter(file)
			if err = writer.Write(summary); err != nil {
				log.Fatalln("error writing to a file!")
			}
			id++
			writer.Flush()
			_ = file.Close()
		}
		return
	}
	res = append(res, cards[index])
	generateAll(cards, index+1, res, id)
	res = res[:len(res)-1]
	generateAll(cards, index+1, res, id)
}

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
	//cardsInFile := random.Intn(7) + 10 // [10, 17]
	cards := make([]card.Card, 0)
	duplicate := make(map[card.Card]int)

	for len(cards) < 52 {
		generatedCard, _ := card.Random(*random)
		if _, found := duplicate[*generatedCard]; !found {
			cards = append(cards, *generatedCard)
		}
		duplicate[*generatedCard]++
	}
	//log.Printf("Generated cards %s\n", cards)
	generateAll(cards, 0, make([]card.Card, 0), 0)
	//}
}

//cards := make([]card.Card, 0)
//for len(cards) < 52 {
//generatedCard, _ := card.Random(*random)
//if _, ok := used[*generatedCard]; !ok {
//cards = append(cards, *generatedCard)
//}
//used[*generatedCard]++
//}
