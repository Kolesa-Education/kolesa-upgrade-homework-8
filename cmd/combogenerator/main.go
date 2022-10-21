package main

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/solution/combinations"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/solution/utility"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

const DatasetDir = "./dataset"
const ResultDir = "./results"

func main() {

	files, err := ioutil.ReadDir(DatasetDir)

	if err != nil {
		log.Fatalln("cannot read dataset directory")
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		fileName := file.Name()

		go func(filename string) {
			defer wg.Done()
			handleFile(filename)
		}(fileName)
	}

	wg.Wait()

}

func handleFile(fileName string) {
	content, err := os.ReadFile(DatasetDir + "/" + fileName)
	if err != nil {
		log.Fatalln("Content of this file cannot read", err)
	}

	cards, err := utility.ConvertCards(strings.TrimSpace(string(content)))
	if err != nil {
		log.Fatalln(err)
	}

	var output [5]card.Card
	length := len(cards)

	combinations.Combinations(cards, output, 0, length-1, 0, 5, ResultDir+"/"+fileName)
}
