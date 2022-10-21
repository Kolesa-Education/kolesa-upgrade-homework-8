package main

import (
	"encoding/csv"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	combinations "github.com/mxschmitt/golang-combinations"
	"github.com/samber/lo"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
)

type Permutations struct {
	cards []card.Card
	title string
}

type OutCSV struct {
	path        string
	resultCards string
}

func cardsData(path string) ([]string, error) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	source := string(input)
	cards := deleteRepeats(strings.Split(source[:len(source)-1], ","))
	return cards, nil
}

func cards(str []string) []card.Card {
	var cards []card.Card
	var cd card.Card
	for _, fa := range str {
		for i, su := range fa {
			if isValidSuit(string(su)) {
				cd.Suit = string(su)
			} else {
				cd.Face = fa[i:]
				break
			}
		}
		cards = append(cards, cd)
	}
	return cards
}

func isValidSuit(suit string) bool {
	switch suit {
	case card.SuitClubsUnicode, card.SuitDiamondsUnicode, card.SuitHeartsUnicode, card.SuitSpadesUnicode:
		return true
	default:
		return false
	}
}

func deleteRepeats(slice []string) []string {
	tmp := make(map[string]bool)
	var processed []string
	for _, val := range slice {
		if _, ok := tmp[val]; !ok {
			tmp[val] = true
			processed = append(processed, val)
		}
	}
	return processed
}
func (p Permutations) getTitle() string {
	return p.title
}
func (o OutCSV) getPath() string {
	return o.path
}
func (o OutCSV) getResult() string {
	return o.resultCards
}
func (p *Permutations) checkPermutation() int {
	if len(p.cards) != 5 {
		return -1
	}

	switch {
	case isPair(p.cards):
		p.title = "Pair"
	case isTwoPairs(p.cards):
		p.title = "Two Pairs"
	case isThreeKind(p.cards):
		p.title = "Three of a Kind"
	case isStraight(p.cards):
		p.title = "Straight"
	case isFlush(p.cards):
		p.title = "Flush"
	case isFullHouse(p.cards):
		p.title = "Full House"
	case isFourKind(p.cards):
		p.title = "Four of a kind"
	case isStraightFlush(p.cards):
		p.title = "Straight Flush"
	default:
		return -1
	}
	return 1
}

func faces(cards []card.Card) map[string]int {
	result := make(map[string]int)
	for _, val := range cards {
		result[val.Face] += 1
	}
	return result
}

func suits(cards []card.Card) map[string]int {
	result := make(map[string]int)
	for _, val := range cards {
		result[val.Suit] += 1
	}
	return result
}

func sortSlice(m map[string]int) []int {
	sorted := make([]int, 0, len(m))
	for key := range m {
		sorted = append(sorted, m[key])
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j]
	})
	return sorted
}

func facePerm(cards []card.Card, perm1, perm2 int) bool {
	faces := sortSlice(faces(cards))
	if faces[0] == perm1 && faces[1] == perm2 {
		return true
	}
	return false
}

func isPair(cards []card.Card) bool {
	return facePerm(cards, 2, 1)
}

func isTwoPairs(cards []card.Card) bool {
	return facePerm(cards, 2, 2)
}

func isThreeKind(cards []card.Card) bool {
	return facePerm(cards, 3, 1)
}
func isStraight(cards []card.Card) bool {
	var cardsFaces []int
	for _, val := range cards {
		cardsFaces = append(cardsFaces, faceEqual(val.Face))
	}
	sort.Slice(cardsFaces, func(i, j int) bool {
		return cardsFaces[i] < cardsFaces[j]
	})
	for i := range cardsFaces {
		if i == 0 && cardsFaces[i] == 0 && cardsFaces[i+1] == 11 {
			continue
		}
		if i != len(cardsFaces)-1 && cardsFaces[i]+1 != cardsFaces[i+1] {
			return false
		}
	}
	return true
}

func isFlush(cards []card.Card) bool {
	suit := suits(cards)
	if len(suit) == 1 {
		return true
	}
	return false
}

func isFullHouse(cards []card.Card) bool {
	return facePerm(cards, 3, 2)
}

func isFourKind(cards []card.Card) bool {
	return facePerm(cards, 4, 1)
}

func isStraightFlush(cards []card.Card) bool {
	if isFlush(cards) && isStraight(cards) {
		return true
	}
	return false
}

func allPermutations(in []string) [][]string {
	result := combinations.Combinations(in, 5)
	return result
}

func faceEqual(in string) int {
	switch in {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	default:
		return -1
	}
}

func cardsString(cards []card.Card) string {
	var result string
	for i := range cards {
		result += fmt.Sprintf("%v%v ", cards[i].Suit, cards[i].Face)
	}
	return result
}

func pokerPermutations(path string, wg *sync.WaitGroup) {
	cardsSlice, err := cardsData(path)
	if err != nil {
	}
	allPerm := allPermutations(cardsSlice)
	var outputPerm string
	for _, val := range allPerm {
		cards := cards(val)
		permutation := Permutations{cards: cards}
		err := permutation.checkPermutation()
		if err != 1 {
			continue
		}
		outputPerm += fmt.Sprintf("%v| %v\n", cardsString(cards), permutation.getTitle())
	}
	result := OutCSV{resultCards: outputPerm}
	resPath := fmt.Sprintf("results/data%v", strings.Trim(path, "dataset/"))
	err = ioutil.WriteFile(resPath, []byte(result.getResult()), 0666)
	if err != nil {
		fmt.Println("error writing to a file!", err)
	}
	wg.Done()
}

func cardsToRepresentations(cards []card.Card) []string {
	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
		r, _ := c.ShortRepresentation()
		return r
	})
	return representations
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
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		path := fmt.Sprintf("dataset/dat%v.csv", i)
		go pokerPermutations(path, &wg)
	}
	wg.Wait()
}
