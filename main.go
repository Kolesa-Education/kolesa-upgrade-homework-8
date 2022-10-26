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
	numPerGoroutine := 1500
	numOfGoroutine := len(cardCombinations) / numPerGoroutine
	reminder := 0
	var wg sync.WaitGroup
	channel := make(chan string, len(cardCombinations))
	if len(cardCombinations)%numPerGoroutine != 0 {
		numOfGoroutine++
		reminder = len(cardCombinations) % numPerGoroutine
	}
	for i := 0; i < numOfGoroutine; i++ {
		rightBound := i*numPerGoroutine + numPerGoroutine
		if reminder != 0 && i == numOfGoroutine-1 {
			rightBound = i*numPerGoroutine + reminder
		}
		wg.Add(1)
		go func(i int, channel chan string) {
			defer wg.Done()
			for j := i * numPerGoroutine; j < rightBound; j++ {

				if len(cardCombinations) == 1 {
					continue
				}

				res := findCombination(cardCombinations[j])
				channel <- res

			}

		}(i, channel)
	}
	wg.Wait()

	for i := 0; i < len(cardCombinations); i++ {
		res := <-channel
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
