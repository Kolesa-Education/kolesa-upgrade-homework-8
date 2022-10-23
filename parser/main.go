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
	var dir = "dataset"
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		waitGroup.Add(1)
		go fileWorker(dir, file.Name(), &waitGroup)
	}
	waitGroup.Wait()
}

// Читает файл
func fileWorker(dir string, file string, waitGroup *sync.WaitGroup) {
	//Чтение файла
	dat, err := os.ReadFile(dir + "/" + file)
	if err != nil {
		println(err.Error())
		return
	}
	//Ищем комбинации
	result := cardParser(string(dat))
	//Если комбинации найдены, записываем в файл
	if result != nil {
		fileWriter(file, result)
	}
	waitGroup.Done()
}

func fileWriter(file string, combinations []map[string][]*card.Card) {
	//Создает файл
	resultFile, err := os.Create("results/" + file)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//Проходит по всем комбинациям
	for _, pockerComb := range combinations {
		for combName, combination := range pockerComb {
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
			_, err2 := resultFile.WriteString(combString + " | " + combName + "\n")
			if err2 != nil {
				log.Println(err2.Error())
				continue
			}
		}
	}
}

// Разбивает каждую карту на масть и значение
func cardParser(pack string) []map[string][]*card.Card {
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
		//println(cardSplitted[0], cardSplitted[1]) //Просто для дебага
		//Создает экземпляр карты
		newCard, err := card.New(cardSplitted[0], cardSplitted[1])
		if err != nil {
			log.Println(err.Error())
			continue
		}
		//Добавляет в массив карт, если в нем еще нет такой же карты
		if !isDuplicate(cardsArr, newCard) {
			cardsArr = append(cardsArr, newCard)
		}
	}
	//Отправляем искать комбинации
	return combinationParser(cardsArr)
}

func combinationParser(cardsArr []*card.Card) []map[string][]*card.Card {
	var pockerComb = make([]map[string][]*card.Card, 0)
	cardCombinations := cardSwitcher(cardsArr)
	if cardCombinations == nil {
		return nil
	}
	//Ищем покерные комбинации для каждой комбинации карт
	for _, cardCombination := range cardCombinations {
		var combName = ""
		if isStraightFlush(cardCombination) {
			combName = "Straight Flush"
			goto APPENDER
		}
		if isFourOfAKind(cardCombination) {
			combName = "Four of a Kind"
			goto APPENDER
		}
		if isFullHouse(cardCombination) {
			combName = "Full house"
			goto APPENDER
		}
		if isFlush(cardCombination) {
			combName = "Flush"
			goto APPENDER
		}
		if isStraight(cardCombination) {
			combName = "Straight"
			goto APPENDER
		}
		if yes, _ := isThreeOfAKind(cardCombination); yes {
			combName = "Three of a Kind"
			goto APPENDER
		}
		if yes, _ := isTwoPairs(cardCombination); yes {
			combName = "Two pairs"
			goto APPENDER
		}
		if yes, _ := isPair(cardCombination); yes {
			combName = "Pair"
			goto APPENDER
		}

	APPENDER:
		//Если комбинация найдена, добавляем ее в массив комбинаций
		println(combName)
		if combName != "" {
			var comb = make(map[string][]*card.Card)
			comb[combName] = cardCombination
			pockerComb = append(pockerComb, comb)
		}
	}
	return pockerComb
}

func isDuplicate(cardsArr []*card.Card, cardToCheck *card.Card) bool {
	//Если это первый элемент массива
	if len(cardsArr) == 0 {
		return false
	}
	//Проверка на дубликат
	for _, card := range cardsArr {
		if card.Suit == cardToCheck.Suit && card.Face == cardToCheck.Face {
			return true
		}
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
	indexesArr := combin.Combinations(len(cardsArr), 5) //[[0,1,2,3,4],[1,2,3,4,5]...]
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

func removeCards(orig []*card.Card, remove []*card.Card) []*card.Card {
	var result = make([]*card.Card, len(orig))
	copy(result, orig)
	//Сравниваем все карты с картами для удаления
	for i := 0; i < len(result); i++ {
		for _, card := range remove {
			if result[i] == card {
				//При совпадении удаляем карту
				result = append(result[:i], result[i+1:]...)
				//Уменьшаем итератор т.к. удалили 1 элемент
				i--
				break
			}
		}
	}
	return result
}

func faceToInt(card *card.Card) int {
	//Маппинг буквенных номиналов
	var faceMap = make(map[string]int)
	faceMap["J"] = 11
	faceMap["Q"] = 12
	faceMap["K"] = 13
	faceMap["A"] = 14
	result, err1 := strconv.Atoi(card.Face)
	//Если номинал буквенный, маппим по карте и возвращаем значение
	if err1 != nil {
		result = faceMap[card.Face]
	}
	return result
}

func isStraightFlush(cards []*card.Card) bool {
	if isFlush(cards) && isStraight(cards) {
		return true
	}
	return false
}

func isFourOfAKind(cards []*card.Card) bool {
	hasTwoPairs, twoPairs := isTwoPairs(cards)
	if !hasTwoPairs {
		return false
	}
	if twoPairs[0].Face == twoPairs[len(twoPairs)-1].Face {
		return true
	}
	return false
}

func isFullHouse(cards []*card.Card) bool {
	hasThree, three := isThreeOfAKind(cards)
	if !hasThree {
		return false
	}
	cards = removeCards(cards, three)
	hasOnePair, _ := isPair(cards)
	if !hasOnePair {
		return false
	}
	return true
}

func isFlush(cards []*card.Card) bool {
	//Проверяем у всех ли карт одинаковая масть
	for i := 1; i < len(cards); i++ {
		if cards[i].Suit != cards[i-1].Suit {
			return false
		}
	}
	return true
}

func isStraight(cards []*card.Card) bool {
	//Сортировка пузырьком
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < len(cards)-i-1; j++ {
			//Перевод номинала в int
			curCard := faceToInt(cards[j])
			nextCard := faceToInt(cards[j+1])
			if curCard > nextCard {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			}
		}
	}
	//Проверка последовательности карт
	for i := 0; i < len(cards)-1; i++ {
		curCard := faceToInt(cards[i])
		nextCard := faceToInt(cards[i+1])
		//Если последовательность нарушена
		if nextCard-curCard != 1 {
			//Проверка на комбинацию A 2 3 4 5
			if nextCard != 14 || faceToInt(cards[0]) != 2 {
				return false
			}
		}
	}
	return true
}

func isThreeOfAKind(cards []*card.Card) (bool, []*card.Card) {
	//Проверяем наличие пары
	hasOnePair, pair := isPair(cards)
	if !hasOnePair {
		return false, nil
	}
	//Удаляем карты, засветившиеся в паре
	cards = removeCards(cards, pair)
	//Если в комбинации пара и тройка
	if cards[0].Face == cards[1].Face && cards[1].Face == cards[2].Face {
		return true, []*card.Card{cards[0], cards[1], cards[2]}
	}
	//Если номинал 1 из оставшихся карт совпадает с номиналом пары, условие выполнено
	for _, card := range cards {
		if pair[0].Face == card.Face {
			return true, append(pair, card)
		}
	}
	return false, nil
}

func isTwoPairs(cards []*card.Card) (bool, []*card.Card) {
	//Ищет первую пару
	hasOnePair, pairs := isPair(cards)
	//Если не нашел - комбинация не удалась
	if !hasOnePair {
		return false, nil
	}
	//copy(cards, removeCards(cards, pairs))
	cards = removeCards(cards, pairs)
	//Проверяем наличие пары, без карты уже имеющейся в паре
	if yes, nextPair := isPair(cards); yes {
		pairs = append(pairs, nextPair[0:]...)
		return true, pairs
	}
	return false, nil
}

func isPair(cards []*card.Card) (bool, []*card.Card) {
	var pair = make([]*card.Card, 0)
	//Сопоставляем каждый элемент массива
	for i, card := range cards {
		for i2, cardToCompare := range cards {
			//Если элементы разные, но совпадает номинал- это пара
			if i != i2 && card.Face == cardToCompare.Face {
				pair = append(pair, cards[i], cards[i2])
				return true, pair
			}
		}
	}
	return false, nil
}
