package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"

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

func addToFile(cards []string, message string, path string) {
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err.Error())
    }

	card := strings.Join(cards[:], ",")
	card += " | " + message + "\n";
	_, err = file.WriteString(card)

	defer file.Close()
}

func findComb(cards []string, path string) {
	freq := make(map[string]int)
	freq2 := make(map[string]int)

	for _ , num :=  range cards {
        temp := ""
		for i := 0; i < len(num); i++ {
			if unicode.IsDigit(rune(num[i])) {
				temp = temp + string(num[i])
			}
		}
		freq[temp] = freq[temp] + 1
    }

	for _ , num :=  range cards {
        temp := ""
		for i := 0; i < len(num); i++ {
			if !unicode.IsDigit(rune(num[i])) {
				temp = temp + string(num[i])
			}
		}
		freq2[temp] = freq2[temp] + 1
    }

	// check for pairs
	for _, v := range freq {
        if v >= 2 {
			addToFile(cards, "Pair", path)
		}
		if v >= 3 {
			addToFile(cards, "Three Of A Kind", path)
		}
		if v >= 4 {
			addToFile(cards, "Four Of A Kind", path)
		}
    }

	for _, v := range freq {
        if v >= 2 {
			addToFile(cards, "Pair", path)
		}
		if v >= 3 {
			addToFile(cards, "Three Of A Kind", path)
		}
		if v >= 4 {
			addToFile(cards, "Four Of A Kind", path)
		}
    }
}

func solve(nums []string, temp []string, k, i int, path string) {
	if (i == len(nums)) {
		return
	}	

	if (k == 0) {
		findComb(temp, path)
		if len(temp) > 0 {
			temp = temp[:len(temp)-1]
		}
		return
	}

	temp = append(temp, nums[i])
	
	solve(nums, temp, k - 1, i + 1, path)
	if len(temp) > 0 {
		temp = temp[:len(temp)-1]
	}
	solve(nums, temp, k, i + 1, path)
}

func combinations(path, data string) {
	arr := strings.Split(data, ",")
	temp := make([]string, 0)

	solve(arr, temp, 5, 0, path)
}

func generate(number int, line string) {
	path := "./results/dat" + strconv.Itoa(number) + ".csv"
	file, err := os.Create(path)     
    if err != nil{                        
        fmt.Println("Unable to create file:", err) 
        os.Exit(1)                          
    }

	combinations(path, line)

    defer file.Close()          
}

func read() {
	path := "results"
	_ = os.Mkdir(path, os.ModePerm)

	for i := 0; i < 100; i++ {
		file := "dataset/dat" + strconv.Itoa(i) + ".csv"
		f, err := os.Open(file)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			generate(i, line)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
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

	read()
}
