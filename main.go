package main

import (
	"fmt"
	"os"
	"strings"

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

func main() {
	f, err := os.Open("./dataset")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	cnt := 0
	for _, file := range files {
		fileName := file.Name()
		dataSlice := getDataFromCSV("dataset/" + fileName)
		dataMap := getCardMapFromSlice(dataSlice)
		dataSlice = getUniqueValuesFromDataMap(dataMap, -1)
		cardCombinations := strings.Split(makeCombinations(dataSlice), ";")
		var combinations []string = make([]string, len(cardCombinations))
		pos := 0
		for i := 0; i < len(cardCombinations); i++ {
			//combination := ""
			dataSlice = strings.Split(cardCombinations[i], ",")
			if len(dataSlice) == 1 {
				continue
			}
			channel := make(chan string)
			go findCombination(dataSlice, channel)
			res := <-channel
			if res == "" {
				continue
			}
			combinations[pos] = res
			pos++
		}

		go writeDataInCSV(combinations, cnt)
		cnt++
	}
}
