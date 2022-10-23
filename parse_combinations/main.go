package main

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"gonum.org/v1/gonum/stat/combin"
	"log"
	"os"
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

// Проверяет на дубли комбинаций с разным расположением карт
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

// Возвращает массив всех возможных комбинаций по 5 карт
func cardSwitcher(cardsArr []*card.Card) [][]*card.Card {
	var cardCombinations = make([][]*card.Card, 0)
	//Проверка на наличие 5 карт
	if len(cardsArr) < 5 {
		log.Println("ERROR: Less than 5 cards!")
		return nil
	}
	//Если карт всего 5, возвращаем их
	if len(cardsArr) == 5 {
		return [][]*card.Card{cardsArr}
	}
	//Получаем все возможные варианты комбинаций индексов массива карт
	indexesArr := combin.Combinations(len(cardsArr), 5)
	//Проходим по каждой комбинации индексов
	for _, indexes := range indexesArr {
		//Заполняем переменную combination элементами, находящимися по полученному индексу
		var combination []*card.Card = nil
		for i := 0; i < len(indexes); i++ {
			index := indexes[i]
			combination = append(combination, cardsArr[index])
		}
		//Добавляем комбинацию в массив комбинаций
		cardCombinations = append(cardCombinations, combination)
	}
	return cardCombinations
}

// Находит все возможные крмбинации из 5 карт
func CombinationFinder(cards []*card.Card, fileName string) {

	if len(cards) < 5 {
		println("no combinations in " + fileName)
		return
	}

	PokerCombinationFinder(cardSwitcher(cards), fileName)

}

// Находит покерные комбинации во всех коминациях карт
func PokerCombinationFinder(combinations [][]*card.Card, fileName string) {
	var (
		sameSuit      = make(map[string][]*card.Card)
		sameFace      = make(map[string][]*card.Card)
		pokerCombsArr []map[string][]*card.Card
	)

	pokerComb := make(map[string][]*card.Card)

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
			pokerCombsArr = append(pokerCombsArr, pokerComb)
			pokerComb = make(map[string][]*card.Card)
			sameFace = make(map[string][]*card.Card)
			sameSuit = make(map[string][]*card.Card)
		}
	}

	println("\n")
	//Пишем в комбинации в фалы
	WriteCombs(fileName, pokerCombsArr)

}

// Пишет комбинации в файлы в директории results
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

// Находит СтритФлеши
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

// Находит сет из 4х карт
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

// Находит ФуллХаус
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

// Находит Флеш
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

// Находит Стрит
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

// Находит сет из трех
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

// Находит две пары
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

// Находит пару
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
