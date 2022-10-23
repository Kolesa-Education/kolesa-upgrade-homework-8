package main

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

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
	cardsParser(string(pack), fileName)
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
func cardsParser(pack string, fileName string) {
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
		newCard, err := card.New(cardsSplit[0], cardsSplit[1])
		if err != nil {
			println(err.Error())
		}
		if cardInPack(newCard, cardsArr) {
			continue
		}
		cardsArr = append(cardsArr, newCard)
	}
	CombinationFinder(cardsArr, fileName)
}

// Проверяет на дубли комбинаций
func combInCombs(combination []*card.Card, combPack [][]*card.Card) bool {
	match := 0
	for _, comb := range combPack {
		for _, c := range comb {
			for i := 0; i <= len(combination)-1; i++ {
				for j := 0; j <= len(comb)-1; j++ {
					if c == combination[i] {
						match++
					}
				}
				if match == 5 {
					return true
				}
			}
		}
		match = 0
	}
	return false
}

// Находит все возможные крмбинации
func CombinationFinder(cards []*card.Card, fileName string) {
	var (
		combinations [][]*card.Card
		combination  []*card.Card
		samecard     bool
	)

	if len(cards) < 5 {
		println("no combinations in " + fileName)
		return
	}

	for i := 0; i <= len(cards)-1; i++ {
		for i1 := 0; i1 <= len(cards)-1; i1++ {
			for i2 := 0; i2 <= len(cards)-1; i2++ {
				for i3 := 0; i3 <= len(cards)-1; i3++ {
					for i4 := 0; i4 <= len(cards)-1; i4++ {
						combination = append(combination, cards[i], cards[i1], cards[i2], cards[i3], cards[i4])
						for j := 0; j < len(combination)-1; j++ {
							for k := j + 1; k <= len(combination)-1; k++ {
								if combination[k] == combination[j] {
									combination = nil
									samecard = true
								}
							}
						}
						if samecard {
							samecard = false
							continue
						}
						if !combInCombs(combination, combinations) {
							combinations = append(combinations, combination)
						}
						combination = nil
					}
				}
			}
		}
	}

	PokerCombinationFinder(combinations, fileName)

}

func PokerCombinationFinder(combinations [][]*card.Card, fileName string) {
	var (
		sameSuit      = make(map[string][]*card.Card)
		sameFace      = make(map[string][]*card.Card)
		pokerCombsArr []map[string][]*card.Card //[]*PokerCombinations
	)

	pokerComb := make(map[string][]*card.Card) //*PokerCombinations

	for _, comb := range combinations {
		for _, c := range comb {
			suit := c.Suit
			face := c.Face
			sameSuit[suit] = append(sameSuit[suit], c)
			sameFace[face] = append(sameFace[face], c)
		}
		if StraightFlush(sameSuit) != "" {
			pokerComb[StraightFlush(sameSuit)] = comb
		} else if FourOfKind(sameFace) != "" {
			pokerComb[FourOfKind(sameFace)] = comb
		} else if FullHouse(sameFace) != "" {
			pokerComb[FullHouse(sameFace)] = comb
		} else if Flush(sameSuit) != "" {
			pokerComb[Flush(sameSuit)] = comb
		} else if Straight(sameFace) != "" {
			pokerComb[Straight(sameFace)] = comb
		} else if ThreeOfKind(sameFace) != "" {
			pokerComb[ThreeOfKind(sameFace)] = comb
		} else if TwoPairs(sameFace) != "" {
			pokerComb[TwoPairs(sameFace)] = comb
		} else if Pair(sameFace) != "" {
			pokerComb[Pair(sameFace)] = comb
		}
		if _, v := pokerComb[""]; !v {
			for _, pc := range pokerCombsArr {
				if reflect.DeepEqual(pc, pokerComb) {
					continue
				}
			}
			pokerCombsArr = append(pokerCombsArr, pokerComb)
			pokerComb = make(map[string][]*card.Card)
		}
	}

	println("\n")

	WriteCombs(fileName, pokerCombsArr)

}

func WriteCombs(name string, pokerCombinations []map[string][]*card.Card) {

	file, err := os.Create("parse_combinations/results/" + name)
	if err != nil {
		println("Unable to create file:", err)
	}

	s := ""

	for _, pokerComb := range pokerCombinations {
		for comb, cards := range pokerComb {
			for i := 0; i <= len(cards)-1; i++ {
				suite, _ := cards[i].SuitUnicode()
				if i < len(cards)-1 {
					s += suite + cards[i].Face + ", "
				} else {
					s += suite + cards[i].Face
				}
			}
			s += " | " + comb + "\n"
		}
	}
	file.WriteString(s)
	file.Close()
}

func StraightFlush(suits map[string][]*card.Card) string {
	s := "straight flush"
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
		return s
	}
	println("no straight flush!")
	return ""
}

func FourOfKind(faces map[string][]*card.Card) string {
	s := "four of a kind!"
	for _, face := range faces {
		if len(face) < 4 {
			continue
		}
		return s
	}
	println("no four of a kind")
	return ""
}

func FullHouse(faces map[string][]*card.Card) string {
	s := "fullhouse"
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
		println("no fullhouse!")
		return ""
	}
	println(s)
	return s
}

func Flush(suits map[string][]*card.Card) string {
	s := "flush"
	for _, suit := range suits {
		if len(suit) < 5 {
			continue
		}
		println(s)
		return s
	}
	println("no flush")
	return ""
}

func Straight(faces map[string][]*card.Card) string {
	s := "straight"
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
		println("no straight!")
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
		println("no straight!")
		return ""
	}
	println(s)
	return s
}

func ThreeOfKind(faces map[string][]*card.Card) string {
	s := "three of a kind"
	for _, face := range faces {
		if len(face) < 3 {
			continue
		}
		println(s)
		return s
	}
	println("no three of a kind!")
	return ""
}

func TwoPairs(faces map[string][]*card.Card) string {
	s := "two pairs"
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
	println("no two pairs!")
	return ""
}

func Pair(faces map[string][]*card.Card) string {
	s := "pair"
	for _, face := range faces {
		if len(face) < 2 {
			continue
		}
		println(s)
		return s
	}
	println("no pairs!")
	return ""
}
