package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
	"unicode/utf8"

	"gonum.org/v1/gonum/stat/combin"
)

var orders = map[string]int{"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5, "8": 6, "9": 7, "10": 8, "J": 9, "Q": 10, "K": 11, "A": 12}

func main() {
	directory, err := ioutil.ReadDir("dataset/")
	if err != nil {
		panic(err)
	}

	for i, file := range directory {
		go makePoker(file.Name(), i)
	}
	time.Sleep(3 * time.Second)
}

func makePoker(filename string, fileCount int) {
	f, err := os.Open(fmt.Sprintf("dataset/%s", filename))
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	var newRecords []string

	for i, record := range records[0] {
		original := true
		for j := i + 1; j < len(records[0]); j++ {
			if record == records[0][j] {
				original = false

			}
		}
		if original {
			newRecords = append(newRecords, records[0][i])
		}

	}

	list := combin.Combinations(len(newRecords), 5)
	combins := Binomial(len(newRecords), 5)
	combinations := make([][]string, combins)
	for i, comb := range list {
		temp := make([]string, 5)
		for j, card := range comb {
			temp[j] = newRecords[card]
		}
		combinations[i] = temp
	}

	for i, card := range combinations {
		check := isPoker(sorting(card))
		if check != 0 {
			combinations[i] = append([]string{strconv.Itoa(check)}, combinations[i]...)
		}
	}
	var poker [][]string
	for _, card := range combinations {
		if len(card) == 6 {
			poker = append(poker, card)
		}
	}

	for i, _ := range poker {
		for j := 0; j < len(poker)-i-1; j++ {
			if string(poker[j][0]) > string(poker[j+1][0]) {
				poker[j], poker[j+1] = poker[j+1], poker[j]
			}
		}
	}
	for i, card := range poker {
		poker[i] = card[1:]
	}

	file, err := os.Create(fmt.Sprintf("results/data%d.csv", fileCount))

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	writer := csv.NewWriter(file)

	for _, card := range poker {
		if err = writer.Write(card); err != nil {
			log.Fatalln("error writing to a file!")
		}
	}

	writer.Flush()
	_ = file.Close()
}

func Binomial(n, k int) int {
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

func isPoker(cards []string) int {
	flush := 0
	firstCover, _ := utf8.DecodeRuneInString(cards[0])
	for _, card := range cards {
		cardCover, _ := utf8.DecodeRuneInString(card)
		if cardCover == firstCover {
			flush += 1

		}
	}
	straight := 1
	for i := 0; i < len(cards)-1; i++ {
		if orders[string(cards[i][3:])]+1 == orders[string(cards[i+1][3:])] {
			straight += 1
		}
	}
	if straight == 4 && string(cards[0][3:]) == "2" && string(cards[4][3:]) == "A" {
		straight += 1
	}
	if flush == 5 && straight == 5 {
		return 1
	}
	if straight == 5 {
		return 5
	}
	if flush == 5 {
		return 4
	}
	nominal := make([]int, 5)
	set := false
	pair := 0
	for i, card := range cards {
		nominal[i] = 0
		for _, card2 := range cards {
			if string(card[3:]) == string(card2[3:]) {
				nominal[i] += 1
			}
		}
		if nominal[i] == 4 {
			return 2
		}
		if nominal[i] == 2 {
			pair += 1
		}
		if nominal[i] == 3 {
			set = true
		}
	}
	if pair == 2 && set {
		return 3
	}
	if set {
		return 6
	}
	if pair == 4 {
		return 7
	}
	if pair == 2 {
		return 8
	}
	return 0
}

func sorting(cards []string) []string {
	for i := 0; i < len(cards); i++ {
		for j := 0; j < len(cards)-i-1; j++ {
			if orders[string(cards[j][3:])] > orders[string(cards[j+1][3:])] {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			}
		}
	}
	return cards
}
