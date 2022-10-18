package main

import (
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
