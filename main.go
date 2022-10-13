package main

import (
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	randomSource := rand.NewSource(seed)
	random := rand.New(randomSource)
	fmt.Printf("Initialized random with seed %d\n", seed)

	fmt.Println("Starting to generate cards...")
	for i := 0; i < 100_000; i++ {
		cardsInFile := random.Intn(7) + 3 // [3, 10)
		cards := make([]card.Card, cardsInFile)

		for j := 0; j < cardsInFile; j++ {
			generatedCard, _ := card.Random(*random)
			cards = append(cards, *generatedCard)
		}

	}
}
