package main

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	files, _ := os.ReadDir("dataset")
	for _, file := range files {
		waitGroup.Add(1)
		fileParser("dataset/"+file.Name(), &waitGroup)
	}
	waitGroup.Wait()
}

// Читает файл
func fileParser(file string, waitGroup *sync.WaitGroup) {
	dat, err := os.ReadFile(file)
	if err != nil {
		println(err.Error())
	}
	cardParser(string(dat))
	waitGroup.Done()
}

// Разбивает каждую карту на масть и значение
func cardParser(pack string) {
	pack = strings.Replace(pack, "\n", "", -1)
	cards := strings.Split(pack, ",")
	var cardsArr = make([]*card.Card, 0)
	for _, cardText := range cards {
		cardSplitted := strings.Split(cardText, "")
		if len(cardSplitted) >= 3 {
			cardSplitted[1] = cardSplitted[1] + cardSplitted[2]
			cardSplitted = cardSplitted[0:2]
		}
		switch cardSplitted[0] {
		case "\u2666":
			cardSplitted[0] = "diamonds"
		case "\u2663":
			cardSplitted[0] = "clubs"
		case "\u2665":
			cardSplitted[0] = "hearts"
		case "\u2660":
			cardSplitted[0] = "spades"
		}
		println(cardSplitted[0], cardSplitted[1])
		newCard, err := card.New(cardSplitted[0], cardSplitted[1])
		if err != nil {
			println(err.Error())
			return
		}
		cardsArr = append(cardsArr, newCard)
	}
	combinationFinder(cardsArr)
}

func combinationFinder(cards []*card.Card) {
	var (
		sameSuit = make(map[string][]*card.Card)
		//containsSuit = make(map[string]bool)
		sameFace = make(map[string][]*card.Card)
		//containsFace = make(map[string]bool)
	)

	if len(cards) < 5 {
		println("No combination!")
		return
	}
	for _, c := range cards {
		suit := c.Suit
		face := c.Face
		sameSuit[suit] = append(sameSuit[suit], c)
		sameFace[face] = append(sameFace[face], c)
	}
	isStraightFlush(sameSuit)
	isFourOfAKind(sameFace)
	isFullHouse(sameFace)
	isFlush(sameSuit)
	isStraight(sameFace)
	isThreeOfAKind(sameFace)
	isTwoPairs(sameFace)
	isPair(sameFace)
	println("\n")
}

func isStraightFlush(suits map[string][]*card.Card) []*card.Card {
	for _, suit := range suits {
		if len(suit) < 5 {
			continue
		}
		var convertedFaces = make(map[int]*card.Card)
		for _, card := range suit {
			switch card.Face {
			case "J":
				convertedFaces[11] = card
			case "Q":
				convertedFaces[12] = card
			case "K":
				convertedFaces[13] = card
			case "A":
				convertedFaces[14] = card
			default:
				var faceInt, _ = strconv.Atoi(card.Face)
				convertedFaces[faceInt] = card
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
		var result = make([]*card.Card, 0)
		for _, c := range comb {
			result = append(result, c)
		}
		println("STRAIGHT FLUSH!",
			result[0].Suit, result[0].Face,
			result[1].Suit, result[1].Face,
			result[2].Suit, result[2].Face,
			result[3].Suit, result[3].Face,
			result[4].Suit, result[4].Face)
		return result
	}
	println("NO STRAIGHT FLUSH!")
	return nil
}

func isFourOfAKind(faces map[string][]*card.Card) []*card.Card {
	for _, face := range faces {
		if len(face) < 4 {
			continue
		}
		println("Four of a Kind!",
			face[0].Suit, face[0].Face,
			face[1].Suit, face[1].Face,
			face[2].Suit, face[2].Face,
			face[3].Suit, face[3].Face)
		return []*card.Card{face[0], face[1], face[2], face[3]}
	}
	println("No Four of a Kind")
	return nil
}

func isFullHouse(faces map[string][]*card.Card) []*card.Card {
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
		return nil
	}
	println("FULLHOUSE!",
		result[0].Suit, result[0].Face,
		result[1].Suit, result[1].Face,
		result[2].Suit, result[2].Face,
		result[3].Suit, result[3].Face,
		result[4].Suit, result[4].Face)
	return result
}

func isFlush(suits map[string][]*card.Card) []*card.Card {
	for _, suit := range suits {
		if len(suit) < 5 {
			continue
		}
		println("FLUSH!",
			suit[0].Suit, suit[0].Face,
			suit[1].Suit, suit[1].Face,
			suit[2].Suit, suit[2].Face,
			suit[3].Suit, suit[3].Face,
			suit[4].Suit, suit[4].Face)
		return []*card.Card{suit[0], suit[1], suit[2], suit[3], suit[4]}
	}
	println("NO FLUSH!")
	return nil
}

func isStraight(faces map[string][]*card.Card) []*card.Card {
	var convertedFaces = make(map[int]*card.Card)
	for _, face := range faces {
		for _, card := range face {
			switch card.Face {
			case "J":
				convertedFaces[11] = card
			case "Q":
				convertedFaces[12] = card
			case "K":
				convertedFaces[13] = card
			case "A":
				convertedFaces[14] = card
			default:
				var faceInt, _ = strconv.Atoi(card.Face)
				convertedFaces[faceInt] = card
			}
		}
	}
	if len(convertedFaces) < 5 {
		println("NO STRAIGHT")
		return nil
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
		return nil
	}
	println("STRAIGHT!",
		result[0].Suit, result[0].Face,
		result[1].Suit, result[1].Face,
		result[2].Suit, result[2].Face,
		result[3].Suit, result[3].Face,
		result[4].Suit, result[4].Face)
	return result
}

func isThreeOfAKind(faces map[string][]*card.Card) []*card.Card {
	for _, face := range faces {
		if len(face) < 3 {
			continue
		}
		println("Three of a Kind!",
			face[0].Suit, face[0].Face,
			face[1].Suit, face[1].Face,
			face[2].Suit, face[2].Face)
		return []*card.Card{face[0], face[1], face[2]}
	}
	println("No Three of a Kind")
	return nil
}

func isTwoPairs(faces map[string][]*card.Card) []*card.Card {
	var pairs []*card.Card
	for _, face := range faces {
		if len(face) >= 4 {
			println("TWO PAIRS!",
				face[0].Suit, face[0].Face,
				face[1].Suit, face[1].Face,
				face[2].Suit, face[2].Face,
				face[3].Suit, face[3].Face)
			return []*card.Card{face[0], face[1], face[2], face[3]}
		}
		if len(face) >= 2 {
			pairs = append(pairs, face[0], face[1])
			if len(pairs) >= 4 {
				println("TWO PAIRS!",
					pairs[0].Suit, pairs[0].Face,
					pairs[1].Suit, pairs[1].Face,
					pairs[2].Suit, pairs[2].Face,
					pairs[3].Suit, pairs[3].Face)
				return pairs
			}
			continue
		}
	}
	println("No Two Pairs")
	return nil
}

func isPair(faces map[string][]*card.Card) []*card.Card {
	for _, face := range faces {
		if len(face) < 2 {
			continue
		}
		println("PAIR!", face[0].Suit, face[0].Face, face[1].Suit, face[1].Face)
		return []*card.Card{face[0], face[1]}
	}
	println("No pairs")
	return nil
}
