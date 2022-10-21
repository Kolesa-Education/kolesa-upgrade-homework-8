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

		go func(filename string) { // Здесь начинаем новую горутину для обработку файлов
			defer wg.Done()
			handleFile(filename)
		}(fileName)
	}

	wg.Wait()
}

func handleFile(fileName string) {
	content, err := os.ReadFile(DatasetDir + "/" + fileName)
	if err != nil {
		log.Printf("content of this %s file cannot read: %s\n", fileName, err.Error())
		return
	}

	//Конвертируем контент файла в структуры карты
	cards, err := utility.ConvertToCards(strings.TrimSpace(string(content)))

	//Убираем дубликаты карт
	cards = utility.UniqueCards(cards)
	if err != nil {
		log.Printf("cannot convert to cards: %s\n", err.Error())
		return
	}

	var output [5]card.Card
	length := len(cards)

	//Находить все комбинаций и обрабатывает каждую из них
	combinations.Combinations(cards, output, 0, length-1, 0, 5, ResultDir+"/"+fileName)
}
