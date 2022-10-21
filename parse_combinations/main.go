package main

import (
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"os"
	"strconv"
	"strings"
	"sync"
)

type PokerCombinations struct {
	cardsList   []*card.Card
	combination string
}

func main() {
	var waitGroup sync.WaitGroup
	files, _ := os.ReadDir("dataset")
	for _, file := range files {
		name := file.Name()
		waitGroup.Add(1)
		go fileParser("dataset/"+file.Name(), &waitGroup, name)
	}
	waitGroup.Wait()
}

// Читает файл
func fileParser(file string, waitGroup *sync.WaitGroup, fileName string) {
	pack, err := os.ReadFile(file)
	if err != nil {
		println(err.Error())
	}
	combs := cardsParser(string(pack))
	WriteCombs(fileName, combs)
	waitGroup.Done()
}

// Проверяет на дубли карт в файле
func cardInPack(card *card.Card, pack []*card.Card) bool {
	for _, b := range pack {
		if b == card {
			return true
		}
	}
	return false
}

// Разбивает каждую карту на масть и значение
func cardsParser(pack string) []*PokerCombinations {
	pack = strings.Replace(pack, "\n", "", -1)
	cards := strings.Split(pack, ",")
	var cardsArr = make([]*card.Card, 0)
	for _, cardText := range cards {
		cardsSplit := strings.Split(cardText, "")
		if len(cardsSplit) >= 3 {
			cardsSplit[1] = cardsSplit[1] + cardsSplit[2]
			cardsSplit = cardsSplit[0:2]
		}
		switch cardsSplit[0] {
		case "\u2666":
			cardsSplit[0] = "diamonds"
		case "\u2663":
			cardsSplit[0] = "clubs"
		case "\u2665":
			cardsSplit[0] = "hearts"
		case "\u2660":
			cardsSplit[0] = "spades"
		}
		println(cardsSplit[0], cardsSplit[1])
		newCard, err := card.New(cardsSplit[0], cardsSplit[1])
		if err != nil {
			println(err.Error())
			return nil
		}
		if cardInPack(newCard, cardsArr) {
			continue
		}
		cardsArr = append(cardsArr, newCard)
	}
	return combinationFinder(cardsArr)
}

// Проверяет на дубли комбинаций
func combInCombs(combination []*card.Card, combPack [][]*card.Card) bool {
	match := 0
	for _, comb := range combPack {
		for _, c := range comb {
			for i := 0; i <= len(combination)-1; i++ {
				if c == combination[i] {
					match++
					if match == 5 {
						return true
					}
				}
			}
		}
		match = 0
	}
	return false
}

// Находит все возможные крмбинации
func combinationFinder(cards []*card.Card) []*PokerCombinations {
	var (
		sameSuit      = make(map[string][]*card.Card)
		sameFace      = make(map[string][]*card.Card)
		combinations  [][]*card.Card
		combination   []*card.Card
		pokerCombsArr []*PokerCombinations
		pokerComb     *PokerCombinations
	)

	if len(cards) < 5 {
		println("No combination!")
		return nil
	}

	for i := 0; i <= len(cards)-1; i++ {
		for i1 := 0; i1 <= len(cards)-1; i1++ {
			for i2 := 0; i2 <= len(cards)-1; i2++ {
				for i3 := 0; i3 <= len(cards)-1; i3++ {
					for i4 := 0; i4 <= len(cards)-1; i4++ {
						combination = append(combination, cards[i], cards[i1], cards[i2], cards[i3], cards[i4])
						if combInCombs(combination, combinations) {
							continue
						}
						combinations = append(combinations, combination)
					}
				}
			}
		}
	}

	for _, comb := range combinations {
		pokerComb.cardsList = comb
		for _, c := range comb {
			suit := c.Suit
			face := c.Face
			sameSuit[suit] = append(sameSuit[suit], c)
			sameFace[face] = append(sameFace[face], c)
		}
		if StraightFlush(sameSuit) != "" {
			pokerComb.combination = StraightFlush(sameSuit)
		} else if FourOfKind(sameFace) != "" {
			pokerComb.combination = FourOfKind(sameFace)
		} else if FullHouse(sameFace) != "" {
			pokerComb.combination = FullHouse(sameFace)
		} else if Flush(sameSuit) != "" {
			pokerComb.combination = Flush(sameSuit)
		} else if Straight(sameFace) != "" {
			pokerComb.combination = Straight(sameFace)
		} else if ThreeOfKind(sameFace) != "" {
			pokerComb.combination = ThreeOfKind(sameFace)
		} else if TwoPairs(sameFace) != "" {
			pokerComb.combination = TwoPairs(sameFace)
		} else if Pair(sameFace) != "" {
			pokerComb.combination = Pair(sameFace)
		}
		if pokerComb.combination != "" {
			pokerCombsArr = append(pokerCombsArr, pokerComb)
		}
	}

	println("\n")
	return pokerCombsArr
}

func WriteCombs(name string, pokerCombinations []*PokerCombinations) {
	file, err := os.Create("results/" + name)
	if err != nil {
		fmt.Println("Unable to create file:", err)
	}

	s := ""

	for _, pokerComb := range pokerCombinations {
		for _, c := range pokerComb.cardsList {
			suite, _ := c.SuitUnicode()
			s += suite + c.Face + ","
		}
		s += " | " + pokerComb.combination + "\n"
	}
	file.WriteString(s)
	file.Close()
}

func StraightFlush(suits map[string][]*card.Card) string {
	s := "STRAIGHT FLUSH!"
	for _, suit := range suits {
		if len(suit) < 5 {
			continue
		}
		var convertedFaces = make(map[int]*card.Card)
		for _, c := range suit {
			switch c.Face {
			case "J":
				convertedFaces[11] = c
			case "Q":
				convertedFaces[12] = c
			case "K":
				convertedFaces[13] = c
			case "A":
				convertedFaces[14] = c
			default:
				var faceInt, _ = strconv.Atoi(c.Face)
				convertedFaces[faceInt] = c
			}
		}

		if len(convertedFaces) < 5 {
			continue
		}
		var comb = make(map[int]*card.Card, 0)
		for i := 2; i < 15; i++ {
			if len(comb) == 5 {
				break
			}
			if _, ok := convertedFaces[2]; len(comb) == 4 && ok {
				if _, ok1 := convertedFaces[14]; ok1 {
					comb[14] = convertedFaces[14]
				}
			}

			if cur, ok := convertedFaces[i]; ok {
				if len(comb) == 0 {
					comb[i] = cur
					continue
				}
				if _, ok1 := comb[i-1]; ok1 {
					comb[i] = cur
				} else {
					comb = make(map[int]*card.Card, 0)
					comb[i] = cur
					continue
				}
			}
		}
		if len(comb) < 5 {
			continue
		}

		println(s)

		return s
	}
	println("NO STRAIGHT FLUSH!")
	return ""
}

func FourOfKind(faces map[string][]*card.Card) string {
	s := "Four of a Kind!"
	for _, face := range faces {
		if len(face) < 4 {
			continue
		}
		println(s)
		return s
	}
	println("No Four of a Kind")
	return ""
}

func FullHouse(faces map[string][]*card.Card) string {
	s := "FULLHOUSE!"
	var result []*card.Card
	for _, face := range faces {
		if len(face) < 2 {
			continue
		}
		if len(face) >= 3 {
			switch len(result) {
			case 0, 2:
				result = append(result, face[0], face[1], face[2])
			case 3:
				result = append(result, face[0], face[1])
			}
			continue
		} else if len(face) >= 2 && (len(result) == 3 || len(result) == 0) {
			result = append(result, face[0], face[1])
			continue
		}
	}
	if len(result) < 5 {
		println("NO FULLHOUSE!")
		return ""
	}
	println(s)
	return s
}

func Flush(suits map[string][]*card.Card) string {
	s := "FLUSH!"
	for _, suit := range suits {
		if len(suit) < 5 {
			continue
		}
		println(s)
		return s
	}
	println("NO FLUSH!")
	return ""
}

func Straight(faces map[string][]*card.Card) string {
	s := "STRAIGHT!"
	var convertedFaces = make(map[int]*card.Card)
	for _, face := range faces {
		for _, c := range face {
			switch c.Face {
			case "J":
				convertedFaces[11] = c
			case "Q":
				convertedFaces[12] = c
			case "K":
				convertedFaces[13] = c
			case "A":
				convertedFaces[14] = c
			default:
				var faceInt, _ = strconv.Atoi(c.Face)
				convertedFaces[faceInt] = c
			}
		}
	}
	if len(convertedFaces) < 5 {
		println("NO STRAIGHT")
		return ""
	}
	var comb = make(map[int]*card.Card)
	for i := 0; i < 15; i++ {
		if len(comb) == 5 {
			break
		}
		if _, ok := convertedFaces[2]; len(comb) == 4 && ok {
			if _, ok1 := convertedFaces[14]; ok1 {
				comb[14] = convertedFaces[14]
				continue
			}
		}

		if cur, ok := convertedFaces[i]; ok {
			if len(comb) == 0 {
				comb[i] = cur
				continue
			}
			if _, ok1 := comb[i-1]; ok1 {
				comb[i] = cur
				continue
			} else {
				comb = make(map[int]*card.Card, 0)
				comb[i] = cur
				continue
			}
		}
	}
	var result = make([]*card.Card, 0)
	for _, c := range comb {
		result = append(result, c)
	}
	if len(result) < 5 {
		println("NO STRAIGHT")
		return ""
	}
	println(s)
	return s
}

func ThreeOfKind(faces map[string][]*card.Card) string {
	s := "Three of a Kind!"
	for _, face := range faces {
		if len(face) < 3 {
			continue
		}
		println(s)
		return s
	}
	println("No Three of a Kind")
	return ""
}

func TwoPairs(faces map[string][]*card.Card) string {
	s := "TWO PAIRS!"
	var pairs []*card.Card
	for _, face := range faces {
		if len(face) >= 4 {
			println(s)
			return s
		}
		if len(face) >= 2 {
			pairs = append(pairs, face[0], face[1])
			if len(pairs) >= 4 {
				println(s)
				return s
			}
			continue
		}
	}
	println("No Two Pairs")
	return ""
}

func Pair(faces map[string][]*card.Card) string {
	s := "PAIR!"
	for _, face := range faces {
		if len(face) < 2 {
			continue
		}
		println(s)
		return s
	}
	println("No pairs")
	return ""
}
