package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
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

func duplicateCount(list []string) map[string]int { // ищет количество одинаковых значении в слайсе

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1
		} else {
			duplicate_frequency[item] = 1
		}
	}
	return duplicate_frequency
}

func checkForValue(cardvalue int, cardsMap map[string]int) bool {

	for _, value := range cardsMap {
		if value == cardvalue {
			return true
		}
	}

	//if value not found return false
	return false
}

func hasDuplicateTwo(m map[string]int) bool {
	x := make(map[int]struct{})

	for _, v := range m {
		if _, has := x[v]; has {
			if v == 2 {
				return true
			}
			return false
		}
		x[v] = struct{}{}
	}

	return false
}

func isFiveConsecutive(faces []int) bool {
	numsMap := make(map[int]int)
	lenFaces := len(faces)

	for i := 0; i < lenFaces; i++ {
		numsMap[faces[i]] = 0
	}

	maxLCS := 0
	for i := 0; i < lenFaces; i++ {
		currentLen := 1
		counter := 1
		for {
			val, ok := numsMap[faces[i]+counter]
			if ok {
				if val != 0 {
					currentLen += val
					break
				} else {
					currentLen += 1
					counter += 1
				}
			} else {
				break
			}
		}

		if currentLen > maxLCS {
			maxLCS = currentLen
		}
		numsMap[faces[i]] = currentLen
	}

	if maxLCS == 5 {
		return true
	} else {
		return false
	}
}

func findPokerCombinations(dupSuit map[string]int, dupFace map[string]int, intFace []int) string { //находим все покерные комбинации
	if checkForValue(3, dupFace) && checkForValue(2, dupFace) { //checkForValue(количество повторении, масть/номинал)
		return "Full House"
	} else if checkForValue(4, dupFace) {
		return "Four of a kind"
	} else if checkForValue(3, dupFace) {
		return "Three of a kind"
	} else if hasDuplicateTwo(dupFace) { //hasDuplicateTwo - если "количество повторении:2" встречается 2 раза)
		return "Two Pairs"
	} else if !hasDuplicateTwo(dupFace) && checkForValue(2, dupFace) {
		return "Pair"
	} else if isFiveConsecutive(intFace) && checkForValue(5, dupSuit) { //isFiveConsecutive - находит есть ли 5 последовательных цифр
		return "Straight Flush"
	} else if isFiveConsecutive(intFace) {
		return "Straight"
	} else if checkForValue(5, dupSuit) {
		return "Flush"
	} else {
		return "noCombinations"
	}
}

func writePokerCombinations(n int) {
	records := readCsvFile(fmt.Sprintf("dataset/dat%d.csv", n))
	distinctrecords := removeDuplicates(records) //удаление дубликатов
	var pokerCombinations []string
	var pokerCombinationsString string

	cs := combin.Combinations(len(distinctrecords), 5) //комбинация из 5 карт
	for _, c := range cs {
		var combinationSuit []string //слайс для мастей
		var combinationFace []string //слайс для номиналов
		var intFace []int            //слайс для цифровых значении номиналов, нужно для последовательного номинала

		for i := 0; i < 5; i++ {
			splitcard := strings.SplitN(distinctrecords[c[i]], "", 2)
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

			intVar, err := strconv.Atoi(splitcard[1])
			if err != nil {
				panic(err)
			}
			intFace = append(intFace, intVar)                       //слайс для цифровых значении номиналов, нужно для последовательного номинала
			combinationSuit = append(combinationSuit, splitcard[0]) //слайс для мастей в комбинации
			combinationFace = append(combinationFace, splitcard[1]) //слайс для номиналов в комбинации
		}

		isFiveConsecutive(intFace)                 //функция для Straight и Straight Flush
		dupSuit := duplicateCount(combinationSuit) //ищем сколько раз повторяются определенные масти в комбинации, получаем map dupSuit [масть: количество повторении]
		dupFace := duplicateCount(combinationFace) //ищем сколько раз повторяются определенные номиналы в комбинации, получаем map dupFace [номинал: количество повторении]

		findPokerCombinations(dupSuit, dupFace, intFace) //функция находит и возвращает названия покерных комбинации

		if findPokerCombinations(dupSuit, dupFace, intFace) != "noCombinations" { //берем все покерные комбинации
			pokerCombination := fmt.Sprintf("%s,%s,%s,%s,%s | %s", distinctrecords[c[0]], distinctrecords[c[1]], distinctrecords[c[2]], distinctrecords[c[3]], distinctrecords[c[4]], findPokerCombinations(dupSuit, dupFace, intFace))
			pokerCombinationsString = pokerCombinationsString + pokerCombination + "\n"
		}

	}

	pokerCombinations = append(pokerCombinations, pokerCombinationsString)
	writeCsvFile(pokerCombinations, n) //записываем все покерные комбинации данного датасета в n-ный csv файл в results
	fmt.Println(pokerCombinations)
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

	ch := make(chan int)
	go channel(ch)
	for v := range ch {
		fmt.Println("done ", v)
	}
}

func channel(chnl chan int) {
	for i := 0; i < 100; i++ {
		writePokerCombinations(i)
		chnl <- i
	}
	close(chnl)
}
