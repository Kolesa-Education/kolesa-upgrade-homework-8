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
		/*go*/ fileParser(file.Name(), &waitGroup)
	}
	waitGroup.Wait()
}

// Читает файл
func fileParser(file string, waitGroup *sync.WaitGroup) {
	dat, err := os.ReadFile("dataset/" + file)
	if err != nil {
		println(err.Error())
		return
	}
	cardParser(string(dat), file)
	waitGroup.Done()
}

// Разбивает каждую карту на масть и значение
func cardParser(pack string, file string) {
	//Убираем все лишнее и разбиваем на отдельные карты
	pack = strings.Replace(pack, "\n", "", -1)
	cards := strings.Split(pack, ",")
	var cardsArr = make([]*card.Card, 0)
	//Проходим по каждому элементу карты
	for _, cardText := range cards {
		cardSplitted := strings.Split(cardText, "") //Дробим по символам
		//Объединяем двузначные значения номинала
		if len(cardSplitted) >= 3 {
			cardSplitted[1] = cardSplitted[1] + cardSplitted[2]
			cardSplitted = cardSplitted[0:2]
		}
		//Меняем значки на текст для передачи в функцию создания карт в будущем
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
		println(cardSplitted[0], cardSplitted[1]) //Просто для дебага
		//Создает экземпляр карты и добавляет в массив карт
		newCard, err := card.New(cardSplitted[0], cardSplitted[1])
		if err != nil {
			println(err.Error())
			return
		}
		cardsArr = append(cardsArr, newCard)
	}
	//Отправляем парситься
	combinations := combinationFinder(cardsArr)
	fileWriter(file, combinations)
}

func combinationFinder(cards []*card.Card) map[string][]*card.Card {
	var ( //Мапы по одинаковым мастям и номиналам
		sameSuit = make(map[string][]*card.Card) //map[масть]карта
		sameFace = make(map[string][]*card.Card) //map[номинал]карта
	)

	//Меньше 5 карт комбинаций нет, так что...
	if len(cards) < 5 {
		println("No combination!")
		return nil
	}
	//Фильтруем по одинаковым мастям и номиналам
	for _, c := range cards {
		suit := c.Suit
		face := c.Face
		sameSuit[suit] = append(sameSuit[suit], c)
		sameFace[face] = append(sameFace[face], c)
	}
	var combinations = make(map[string][]*card.Card)
	//Ищем комбинации
	sf := isStraightFlush(sameSuit)
	if sf != nil {
		combinations["Straight Flush"] = sf
	}
	foak := isFourOfAKind(sameFace)
	if foak != nil {
		combinations["Four Of AKind"] = foak
	}
	fh := isFullHouse(sameFace)
	if fh != nil {
		combinations["Full House"] = fh
	}
	f := isFlush(sameSuit)
	if f != nil {
		combinations["Flush"] = f
	}
	s := isStraight(sameFace)
	if s != nil {
		combinations["Straight"] = s
	}
	tok := isThreeOfAKind(sameFace)
	if tok != nil {
		combinations["Three Of A Kind"] = tok
	}
	tp := isTwoPairs(sameFace)
	if tp != nil {
		combinations["Two Pairs"] = tp
	}
	p := isPair(sameFace)
	if p != nil {
		combinations["Pair"] = p
	}
	println("\n")
	//Возвращаем найденные комбинации
	return combinations
}

func fileWriter(file string, combinations map[string][]*card.Card) {
	//Создает файл
	packFile, err := os.Create("results/" + file)
	if err != nil {
		println(err.Error())
		return
	}
	for combName, combination := range combinations {
		//Генерирует строку комбинации
		var combString = ""
		for _, card := range combination {
			if combString != "" {
				combString += ","
			}
			shortRep, _ := card.ShortRepresentation()
			combString += shortRep
		}
		//Записывает комбинацию в файл
		_, err2 := packFile.WriteString(combString + " | " + combName + "\n")
		if err2 != nil {
			println(err2.Error())
			continue
		}
	}
}

// Просто жесть, переделать если будет время.
func isStraightFlush(suits map[string][]*card.Card) []*card.Card {
	for _, suit := range suits {
		//Нужно только 5 карт одной масти
		if len(suit) < 5 {
			continue
		}
		//Превращаем буквенные номиналы в числовые для упрощения подсчета, попутно отсеиваем дубли
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

		//Если карт без дублей меньше 5 - ищем другие варианты
		if len(convertedFaces) < 5 {
			continue
		}
		//Проверяем комбинацию карт
		var comb = make(map[int]*card.Card, 0)
		for i := 2; i < 15; i++ {
			//Стоп если комбинация уже собрана
			if len(comb) == 5 {
				break
			}
			//Если есть двойка, можем завершить комбинацию тузом
			if _, ok := convertedFaces[2]; len(comb) == 4 && ok {
				//Если есть туз- добавляем его в комбинацию
				if _, ok1 := convertedFaces[14]; ok1 {
					comb[14] = convertedFaces[14]
				}
			}

			//Если номинал есть в списке карт
			if cur, ok := convertedFaces[i]; ok {
				//Если первый элемент- в любом случае добавляем его в комбинацию и ищем дальше
				if len(comb) == 0 {
					comb[i] = cur
					continue
				}
				//Если в комбинации есть номинал на 1 меньше, добавляем карту в комбинацию
				if _, ok1 := comb[i-1]; ok1 {
					comb[i] = cur
				} else { //А если нет- обнуляем комбинацию и записываем первый элемент комбинации
					comb = make(map[int]*card.Card, 0)
					comb[i] = cur
					continue
				}
			}
		}
		//Продолжаем пока не наберется комбинация из 5
		if len(comb) < 5 {
			continue
		}
		//Приводим в годный для возвращения вид и возвращаем
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
		//ищем 4 карты одного номинала
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
		//Меньше 2 карт нам не надо
		if len(face) < 2 {
			continue
		}
		//Если нашли 3 карты одного номинала
		if len(face) >= 3 {
			switch len(result) {
			case 0, 2: //Если еще нет совпадений или пара - добавляем тройку
				result = append(result, face[0], face[1], face[2])
			case 3: //Если уже есть тройка- добавляем пару
				result = append(result, face[0], face[1])
			}
			continue
			//Пару ищем только если ее еще нет
		} else if len(face) >= 2 && (len(result) == 3 || len(result) == 0) {
			result = append(result, face[0], face[1])
			continue
		}
	}
	//Если нашли меньше 5 карт - увы...
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
		//В поисках 5 карт одной масти...
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

// Просто жесть, переделать если будет время.
func isStraight(faces map[string][]*card.Card) []*card.Card {
	var convertedFaces = make(map[int]*card.Card)
	for _, face := range faces {
		//Превращаем буквенные номиналы в числовые для упрощения подсчета, попутно отсеиваем дубли
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
	//Если карт без дублей меньше 5 - ищем другие варианты
	if len(convertedFaces) < 5 {
		println("NO STRAIGHT")
		return nil
	}
	//Проверяем комбинацию карт
	var comb = make(map[int]*card.Card)
	for i := 0; i < 15; i++ {
		//Стоп если комбинация уже собрана
		if len(comb) == 5 {
			break
		}
		//Если есть двойка, можем завершить комбинацию тузом
		if _, ok := convertedFaces[2]; len(comb) == 4 && ok {
			//Если есть туз- добавляем его в комбинацию
			if _, ok1 := convertedFaces[14]; ok1 {
				comb[14] = convertedFaces[14]
				continue
			}
		}
		//Если номинал есть в списке карт
		if cur, ok := convertedFaces[i]; ok {
			//Если первый элемент- в любом случае добавляем его в комбинацию и ищем дальше
			if len(comb) == 0 {
				comb[i] = cur
				continue
			}
			//Если в комбинации есть номинал на 1 меньше, добавляем карту в комбинацию
			if _, ok1 := comb[i-1]; ok1 {
				comb[i] = cur
				continue
			} else { //А если нет- обнуляем комбинацию и записываем первый элемент комбинации
				comb = make(map[int]*card.Card, 0)
				comb[i] = cur
				continue
			}
		}
	}
	//Нет 5 карт в комбинации - нет комбинации
	if len(comb) < 5 {
		println("NO STRAIGHT")
		return nil
	}
	//Приводим в годный для возвращения вид и возвращаем
	var result = make([]*card.Card, 0)
	for _, c := range comb {
		result = append(result, c)
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
		//В поисках 3 карт одного номинала...
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
	//Проходим по каждому номиналу
	for _, face := range faces {
		//Если 4 карты одного номинала- бинго
		if len(face) >= 4 {
			println("TWO PAIRS!",
				face[0].Suit, face[0].Face,
				face[1].Suit, face[1].Face,
				face[2].Suit, face[2].Face,
				face[3].Suit, face[3].Face)
			return []*card.Card{face[0], face[1], face[2], face[3]}
		}
		//Ищем пары
		if len(face) >= 2 {
			pairs = append(pairs, face[0], face[1])
			//Останавливаемся когда найдено 2 пары
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
	//Просто ищем 2 карты с одинаковым номиналом
	for _, face := range faces {
		if len(face) < 2 {
			continue
		}
		//Возвращаем 2 первые карты одинакового номинала
		println("PAIR!", face[0].Suit, face[0].Face, face[1].Suit, face[1].Face)
		return []*card.Card{face[0], face[1]}
	}
	println("No pairs")
	return nil
}
