package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func getCombinationsPerFile(file os.FileInfo) []string {
	fileName := file.Name()
	dataSlice := getDataFromCSV("dataset/" + fileName)
	dataMap := getCardMapFromSlice(dataSlice)
	dataSlice = getUniqueValuesFromDataMap(dataMap, -1)
	cardCombinations := makeCombinationsFromMasks(dataSlice, 5)
	combinations := []string{}
	for _, cardCombination := range cardCombinations {
		if len(cardCombinations) == 1 {
			continue
		}

		res := findCombination(cardCombination)
		if res != "" {
			combinations = append(combinations, res)
		}
	}
	return combinations
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

	start := time.Now()
	var wg sync.WaitGroup
	for i, file := range files {
		wg.Add(1)
		go func(file os.FileInfo, i int) {
			combinations := getCombinationsPerFile(file)
			writeDataInCSV(combinations, i)
			wg.Done()
		}(file, i)

	}
	wg.Wait()

	log.Printf("main, execution time %s\n", time.Since(start))
}
