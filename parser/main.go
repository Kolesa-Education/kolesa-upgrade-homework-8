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

type pockerCombination struct {
	cards []*card.Card
	name  string
}

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

func newCombination(cards []*card.Card, name string) *pockerCombination {
	return &pockerCombination{
		cards: cards,
		name:  name,
	}
}

func combinationToString(comb *pockerCombination) string {
	var result = ""
	for _, crd := range comb.cards {
		if result != "" {
			result += ","
		}
		shortRep, _ := crd.ShortRepresentation()
		result += shortRep
	}
	result += " | " + comb.name
	return result
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

func fileWriter(file string, combinations []*pockerCombination) {
	//Создает файл
	resultFile, err := os.Create("results/" + file)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//Проходит по всем комбинациям и записываем их
	for _, combination := range combinations {
		_, err2 := resultFile.WriteString(combinationToString(combination) + "\n")
		if err2 != nil {
			log.Println(err2.Error())
			continue
		}
	}
}

// Разбивает каждую карту на масть и значение
func cardParser(pack string) []*pockerCombination {
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
	return combinationGenerator(cardsArr)
}

func combinationParser(cardCombination []*card.Card) *pockerCombination {
	if cardCombination == nil {
		return nil
	}
	//Ищем покерные комбинации
	if IsStraightFlush(cardCombination) {
		return newCombination(cardCombination, "Straight Flush")
	}
	if IsFourOfAKind(cardCombination) {
		return newCombination(cardCombination, "Four of a Kind")
	}
	if IsFullHouse(cardCombination) {
		return newCombination(cardCombination, "Full house")
	}
	if IsFlush(cardCombination) {
		return newCombination(cardCombination, "Flush")
	}
	if IsStraight(cardCombination) {
		return newCombination(cardCombination, "Straight")
	}
	if yes, _ := IsThreeOfAKind(cardCombination); yes {
		return newCombination(cardCombination, "Three of a Kind")
	}
	if yes, _ := IsTwoPairs(cardCombination); yes {
		return newCombination(cardCombination, "Two pairs")
	}
	if yes, _ := IsPair(cardCombination); yes {
		return newCombination(cardCombination, "Pair")
	}
	return nil
}

func isDuplicate(cardsArr []*card.Card, cardToCheck *card.Card) bool {
	//Если это первый элемент массива
	if len(cardsArr) == 0 {
		return false
	}
	//Проверка на дубликат
	for _, crd := range cardsArr {
		if crd.Suit == cardToCheck.Suit && crd.Face == cardToCheck.Face {
			return true
		}
	}
	return false
}

// Генерирует все возможные комбинации по 5 карт и их силу
func combinationGenerator(cardsArr []*card.Card) []*pockerCombination {
	if len(cardsArr) < 5 {
		log.Println("ERROR: Less than 5 cards!")
		return nil
	}
	if len(cardsArr) == 5 {
		return []*pockerCombination{combinationParser(cardsArr)}
	}
	var pockerCombinations = make([]*pockerCombination, 0)
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
		//Определяем силу комбинации
		combinationResult := combinationParser(combination)
		if combinationResult == nil {
			continue
		}
		//Добавляем комбинацию в массив комбинаций
		pockerCombinations = append(pockerCombinations, combinationResult)
	}
	return pockerCombinations
}

func removeCards(orig []*card.Card, remove []*card.Card) []*card.Card {
	var result = make([]*card.Card, len(orig))
	copy(result, orig)
	//Сравниваем все карты с картами для удаления
	for i := 0; i < len(result); i++ {
		for _, crd := range remove {
			if result[i] == crd {
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

func IsStraightFlush(cards []*card.Card) bool {
	if IsFlush(cards) && IsStraight(cards) {
		return true
	}
	return false
}

func IsFourOfAKind(cards []*card.Card) bool {
	hasTwoPairs, twoPairs := IsTwoPairs(cards)
	if !hasTwoPairs {
		return false
	}
	if twoPairs[0].Face == twoPairs[len(twoPairs)-1].Face {
		return true
	}
	return false
}

func IsFullHouse(cards []*card.Card) bool {
	hasThree, three := IsThreeOfAKind(cards)
	if !hasThree {
		return false
	}
	cards = removeCards(cards, three)
	hasOnePair, _ := IsPair(cards)
	if !hasOnePair {
		return false
	}
	return true
}

func IsFlush(cards []*card.Card) bool {
	//Проверяем у всех ли карт одинаковая масть
	for i := 1; i < len(cards); i++ {
		if cards[i].Suit != cards[i-1].Suit {
			return false
		}
	}
	return true
}

func IsStraight(cards []*card.Card) bool {
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

func IsThreeOfAKind(cards []*card.Card) (bool, []*card.Card) {
	//Проверяем наличие пары
	hasOnePair, pair := IsPair(cards)
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
	for _, crd := range cards {
		if pair[0].Face == crd.Face {
			return true, append(pair, crd)
		}
	}
	return false, nil
}

func IsTwoPairs(cards []*card.Card) (bool, []*card.Card) {
	//Ищет первую пару
	hasOnePair, pairs := IsPair(cards)
	//Если не нашел - комбинация не удалась
	if !hasOnePair {
		return false, nil
	}
	//copy(cards, removeCards(cards, pairs))
	cards = removeCards(cards, pairs)
	//Проверяем наличие пары, без карты уже имеющейся в паре
	if yes, nextPair := IsPair(cards); yes {
		pairs = append(pairs, nextPair[0:]...)
		return true, pairs
	}
	return false, nil
}

func IsPair(cards []*card.Card) (bool, []*card.Card) {
	var pair = make([]*card.Card, 0)
	//Сопоставляем каждый элемент массива
	for i, crd := range cards {
		for i2, cardToCompare := range cards {
			//Если элементы разные, но совпадает номинал- это пара
			if i != i2 && crd.Face == cardToCompare.Face {
				pair = append(pair, cards[i], cards[i2])
				return true, pair
			}
		}
	}
	return false, nil
}
