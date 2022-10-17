package main

import (
	"encoding/csv"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/ernestosuarez/itertools"
	"github.com/samber/lo"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Наверное по понятиям Golang-a - это лютый говнокод, но я старался

func removeDuplicateElement(arr []string) []string {
	//Т.К. дубликаты мы не учитываем, удаляем их в исходной строке
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func removeMast(arr []string) []string {
	//Удаляем масть, нужно для комбинации Straight
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for _, item := range arr {
		res := strings.ReplaceAll(item, "♦", "")
		res_2 := strings.ReplaceAll(res, "♠", "")
		res_3 := strings.ReplaceAll(res_2, "♥", "")
		res_4 := strings.ReplaceAll(res_3, "♣", "")
		temp[res_4] = struct{}{}
		result = append(result, res_4)
	}
	return result
}
func isStraight(v []string) bool {
	//А вот и она (5 последовательных карт)
	leng := len(v)
	v_mast := removeMast(v)
	for i := 0; i < leng-1; i++ {
		for j := 0; j < leng-1; j++ {
			//Не нашёл другого способа как отсортировать буквы...
			if v_mast[j] == "A" {
				v_mast[j] = "14"
			}
			if v_mast[j] == "K" {
				v_mast[j] = "13"
			}
			if v_mast[j] == "Q" {
				v_mast[j] = "12"
			}
			if v_mast[j] == "J" {
				v_mast[j] = "11"
			}

		}
		for j := 0; j < leng-i-1; j++ {
			first_num, _ := strconv.ParseInt(v_mast[j], 10, 64)
			second_num, _ := strconv.ParseInt(v_mast[j+1], 10, 64)
			if first_num > second_num {
				v_mast[j], v_mast[j+1] = v_mast[j+1], v_mast[j]

			}

		}
	}
	//Здесь я решил захардкодить эти проверки
	first_num, _ := strconv.Atoi(v_mast[0])
	second_num, _ := strconv.Atoi(v_mast[1])
	third_num, _ := strconv.Atoi(v_mast[2])
	fourth_num, _ := strconv.Atoi(v_mast[3])
	fifth_num, _ := strconv.Atoi(v_mast[4])
	//Если след число меньше предыдущего на 1
	if ((second_num - first_num) == 1) && ((third_num - second_num) == 1) && ((fourth_num - third_num) == 1) && ((fifth_num - fourth_num) == 1) {
		return true

	} else /*Проверка на случай комбинации A, 2, 3, 4, 5*/ if fifth_num == 14 && first_num == 2 && second_num == 3 && third_num == 4 && fourth_num == 5 {
		return true
	}
	return false
}
func isFlush(v []string) bool {
	//5 карт одинаковой масти
	if v[0][0:3] == v[1][0:3] && v[1][0:3] == v[2][0:3] && v[2][0:3] == v[3][0:3] && v[3][0:3] == v[4][0:3] {
		return true
	} else {
		return false
	}

}
func isStraightFlush(v []string) bool {
	//Соединил 2 функции
	if isStraight(v) && isFlush(v) {
		return true
	}
	return false
}
func isPair(v []string) bool {
	//Пара
	counter := 0
	for i := 1; i <= len(v); i++ {
		if i == len(v) {
			//Если при первой итерации не нашли пару
			counter += 1
			i = 1 + counter
			if counter == 4 {
				break
			}
		}
		basic_num := v[counter][3:4]
		next_num := v[i][3:4]
		if basic_num == next_num {
			return true
		}
	}
	return false
}
func isThreeOfKind(v []string) bool {
	//Тройка
	counter := 0
	for i := 1; i <= len(v); i++ {
		if i == len(v) {
			counter += 1
			i = 1 + counter
			if counter == 4 {
				break
			}
		}
		basic_num := v[counter][3:4]
		next_num := v[i][3:4]
		if basic_num == next_num {
			trird_num := basic_num
			index_first := counter
			index_second := i

			counter := 0
			for i := 1; i <= len(v); i++ {
				if i == len(v) {
					counter += 1
					i = 1 + counter
					if counter == 4 {
						break
					}
				}
				basic_num := trird_num
				next_num := v[i][3:4]
				if basic_num == next_num && i != index_first && i != index_second {
					return true
				}
			}

			break
		}
	}
	return false
}
func isFourOfKind(v []string) bool {
	//Четвёрка
	counter := 0
	for i := 1; i <= len(v); i++ {
		if i == len(v) {
			counter += 1
			i = 1 + counter
			if counter == 4 {
				break
			}
		}
		basic_num := v[counter][3:4]
		next_num := v[i][3:4]
		if basic_num == next_num {
			first_num := basic_num
			index_first := counter
			index_second := i
			counter := 0
			for i := 1; i <= len(v); i++ {
				if i == len(v) {
					counter += 1
					i = 1 + counter
					if counter == 4 {
						break
					}
				}
				basic_num := first_num
				next_num := v[i][3:4]
				if basic_num == next_num && i != index_first && i != index_second {
					index_third := i
					counter := 0
					for i := 1; i <= len(v); i++ {
						if i == len(v) {
							counter += 1
							i = 1 + counter
							if counter == 4 {
								break
							}
						}
						basic_num := first_num
						next_num := v[i][3:4]
						if basic_num == next_num && i != index_first && i != index_second && i != index_third {
							return true
						}
					}
					break
				}

			}

			break
		}
	}
	return false
}
func isdoublePair(v []string) bool {
	//Две пары
	counter := 0
	for i := 1; i <= len(v); i++ {
		if i == len(v) {

			counter += 1
			i = 1 + counter
			if counter == 4 {
				break
			}
		}
		basic_num := v[counter][3:4]
		next_num := v[i][3:4]
		if basic_num == next_num {
			f_card := basic_num
			counter := 0
			for i := 1; i <= len(v); i++ {
				if i == len(v) {
					//Если при первой итерации не нашли пару
					counter += 1
					i = 1 + counter
					if counter == 4 {
						break
					}
				}
				basic_num := v[counter][3:4]
				next_num := v[i][3:4]
				if basic_num == next_num && basic_num != f_card {
					f_card = basic_num
					return true
				} else {
				}

			}
			break
		}
	}
	return false
}
func isFullHouse(v []string) bool {
	//3 + 2
	counter := 0
	for i := 1; i <= len(v); i++ {
		if i == len(v) {
			counter += 1
			i = 1 + counter
			if counter == 4 {
				break
			}
		}
		basic_num := v[counter][3:4]
		next_num := v[i][3:4]
		if basic_num == next_num {
			trird_num := basic_num
			index_first := counter
			index_second := i
			counter := 0
			for i := 1; i <= len(v); i++ {
				if i == len(v) {
					counter += 1
					i = 1 + counter
					if counter == 4 {
						break
					}
				}
				basic_num := trird_num
				next_num := v[i][3:4]
				if basic_num == next_num && i != index_first && i != index_second {
					first_pair_num := basic_num
					counter := 0
					for i := 1; i <= len(v); i++ {
						if i == len(v) {
							counter += 1
							i = 1 + counter
							if counter == 4 {
								break
							}
						}
						basic_num := v[counter][3:4]
						next_num := v[i][3:4]
						if basic_num == next_num && basic_num != first_pair_num {
							return true
						}
					}
					break
				}
			}

			break
		}
	}
	return false
}
func cardsToRepresentations(cards []card.Card) []string {
	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
		r, _ := c.ShortRepresentation()
		return r
	})
	return representations
}

func logicGo(i int) {
	//Проходимся по всем комбинациям во всех файлах
	var combin string
	file, _ := os.Open(fmt.Sprintf("dataset/dat%d.csv", i))
	b, _ := ioutil.ReadAll(file)
	fmt.Printf("Строка %s\n", string(b))
	_ = file.Close()
	dwarfs := []string{}
	words := strings.Split(string(b), ",")
	for _, word := range words {
		dwarfs = append(dwarfs, word)
	}
	dwarfs = removeDuplicateElement(dwarfs)
	fmt.Println(dwarfs)
	filenew, _ := os.OpenFile(fmt.Sprintf("results/dat%d.csv", i),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	for v := range itertools.CombinationsStr(dwarfs, 5) {
		combin = combin[:0]
		fmt.Println(v)
		if !(isStraightFlush(v)) {
			if !(isFourOfKind(v)) {
				if !(isFullHouse(v)) {
					if !(isFlush(v)) {
						if !(isStraight(v)) {
							if !(isThreeOfKind(v)) {
								if !(isdoublePair(v)) {
									if !(isPair(v)) {
										continue
									} else {
										combin = "PAIR"
									}
								} else {
									combin = "DoublePair"
								}
							} else {
								combin = "ThreeOfKind"
							}
						} else {
							combin = "Straight"
						}
					} else {
						combin = "Flush"
					}
				} else {
					combin = "FullHouse"
				}
			} else {
				combin = "FourOfKind"
			}
		} else {
			combin = "StraightFlush"

		}
		filenew.WriteString("--------------------------------------------\n")
		filenew.WriteString("Набор: " + "\n" + strings.Join(v[:], ",") + "\nКомбинация:" + combin + "\n")
		_ = file.Close()
	}
}
func main() {
	start := time.Now()
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
		go logicGo(i)
		//Добавил задержку, т.к. если выполнять без неё, то большая часть данных будет теряться, из лекции Хамбара я так понял что это норма?
		time.Sleep(5 * time.Second)
		duration := time.Since(start)
		fmt.Println(duration)
	}
}
