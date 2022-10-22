package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/samber/lo"
)

func removeDuplicateElement(cards []string) []string {
	result := make([]string, 0, len(cards))
	temp := map[string]struct{}{}
	for _, item := range cards {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func isPair(hand [5]string) bool {
	for i := 0; i < len(hand)-1; i++ {
		for j := i + 1; j < len(hand); j++ {
			if hand[i][3:] == hand[j][3:] {
				return true
				break
			}
		}
	}
	return false
}

func isTwoPairs(hand [5]string) bool {
	pair := ""
	for i := 0; i < len(hand)-1; i++ {
		for j := i + 1; j < len(hand); j++ {
			if hand[i][3:] == hand[j][3:] && pair == "" {
				pair = hand[i][3:]
				_ = pair
			}
			if hand[i][3:] == hand[j][3:] && pair != hand[i][3:] && pair != "" {
				return true
				break
			}
		}
	}
	return false
}

func isThreeOfAKind(hand [5]string) bool {
	for i := 0; i < len(hand)-3; i++ {
		for j := i + 1; j < len(hand)-2; j++ {
			for k := j + 1; k < len(hand)-1; k++ {
				if hand[i][3:] == hand[j][3:] && hand[j][3:] == hand[k][3:] {
					return true
				}
			}
		}
	}
	return false
}

func isStraight(hand [5]string) bool {
	var clearHand [5]string

	for a := 0; a < 2; a++ {
		for i := 0; i < len(hand); i++ {
			clearHand[i] = hand[i][3:]
			if hand[i][3:] == "J" {
				clearHand[i] = "11"
			}
			if hand[i][3:] == "Q" {
				clearHand[i] = "12"
			}
			if hand[i][3:] == "K" {
				clearHand[i] = "13"
			}
			if hand[i][3:] == "A" {
				if a == 0 {
					clearHand[i] = "14"
				} else {
					clearHand[i] = "1"
				}
			}
		}

		intHand := strToIntArray(clearHand)

		for i := 0; i < len(intHand)-4; i++ {
			for j := i + 1; j < len(intHand)-3; j++ {
				for k := j + 1; k < len(intHand)-2; k++ {
					for l := k + 1; l < len(intHand)-1; l++ {
						for m := l + 1; m < len(intHand); m++ {
							if intHand[i]+1 == intHand[j] && intHand[j]+1 == intHand[k] && intHand[k]+1 == intHand[l] && intHand[l]+1 == intHand[m] {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

func isFlush(hand [5]string) bool {
	if hand[0][0:3] == hand[1][0:3] && hand[1][0:3] == hand[2][0:3] && hand[2][0:3] == hand[3][0:3] && hand[3][0:3] == hand[4][0:3] {
		return true
	}
	return false
}

func isFullHouse(hand [5]string) bool {
	var clearHand [5]string

	for i := 0; i < len(hand); i++ {
		clearHand[i] = hand[i][3:]
		if hand[i][3:] == "J" {
			clearHand[i] = "11"
		}
		if hand[i][3:] == "Q" {
			clearHand[i] = "12"
		}
		if hand[i][3:] == "K" {
			clearHand[i] = "13"
		}
		if hand[i][3:] == "A" {
			clearHand[i] = "14"
		}
	}

	intHand := strToIntArray(clearHand)

	for i := 0; i < len(intHand)-4; i++ {
		for j := i + 1; j < len(intHand)-3; j++ {
			for k := j + 1; k < len(intHand)-2; k++ {
				for l := k + 1; l < len(intHand)-1; l++ {
					for m := l + 1; m < len(intHand); m++ {
						if intHand[i] == intHand[j] && intHand[j] == intHand[k] && intHand[l] == intHand[m] {
							return true
						}
						if intHand[i] == intHand[j] && intHand[k] == intHand[l] && intHand[l] == intHand[m] {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func isFourOfAKind(hand [5]string) bool {
	for i := 0; i < len(hand)-4; i++ {
		for j := i + 1; j < len(hand)-3; j++ {
			for k := j + 1; k < len(hand)-2; k++ {
				for l := k + 1; l < len(hand)-1; l++ {
					for m := l + 1; m < len(hand); m++ {
						if hand[i][3:] == hand[j][3:] && hand[j][3:] == hand[k][3:] && hand[k][3:] == hand[l][3:] {
							return true
						}
						if hand[j][3:] == hand[k][3:] && hand[k][3:] == hand[l][3:] && hand[l][3:] == hand[m][3:] {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func isStraightFlush(hand [5]string) bool {
	var clearHand [5]string

	if hand[0][0:3] == hand[1][0:3] && hand[1][0:3] == hand[2][0:3] && hand[2][0:3] == hand[3][0:3] && hand[3][0:3] == hand[4][0:3] {
		for a := 0; a < 2; a++ {
			for i := 0; i < len(hand); i++ {
				clearHand[i] = hand[i][3:]
				if hand[i][3:] == "J" {
					clearHand[i] = "11"
				}
				if hand[i][3:] == "Q" {
					clearHand[i] = "12"
				}
				if hand[i][3:] == "K" {
					clearHand[i] = "13"
				}
				if hand[i][3:] == "A" {
					if a == 0 {
						clearHand[i] = "14"
					} else {
						clearHand[i] = "1"
					}

				}
			}

			intHand := strToIntArray(clearHand)

			for i := 0; i < len(intHand)-4; i++ {
				for j := i + 1; j < len(intHand)-3; j++ {
					for k := j + 1; k < len(intHand)-2; k++ {
						for l := k + 1; l < len(intHand)-1; l++ {
							for m := l + 1; m < len(intHand); m++ {
								if intHand[i]+1 == intHand[j] && intHand[j]+1 == intHand[k] && intHand[k]+1 == intHand[l] && intHand[l]+1 == intHand[m] {
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func strToIntArray(clearHand [5]string) []int {
	intHand := []int{1, 2, 3, 4, 5}
	intHand[0], _ = strconv.Atoi(clearHand[0])
	intHand[1], _ = strconv.Atoi(clearHand[1])
	intHand[2], _ = strconv.Atoi(clearHand[2])
	intHand[3], _ = strconv.Atoi(clearHand[3])
	intHand[4], _ = strconv.Atoi(clearHand[4])

	sort.Ints(intHand)
	return intHand
}

func cardsToRepresentations(cards []card.Card) []string {
	representations := lo.Map[card.Card, string](cards, func(c card.Card, index int) string {
		r, _ := c.ShortRepresentation()
		return r
	})
	return representations
}

func result(i int) {
	var hand [5]string

	file, _ := os.Open(fmt.Sprintf("dataset/dat%d.csv", i))
	fileData, _ := ioutil.ReadAll(file)
	_ = file.Close()

	editedFileData := strings.Trim(string(fileData), "\n")
	arrayFileData := strings.Split(string(editedFileData), ",")
	fmt.Println(arrayFileData)

	arrayFileData = removeDuplicateElement(arrayFileData)
	fmt.Println(arrayFileData)

	resFile, err := os.Create(fmt.Sprintf("results/dat%d.csv", i))
	_ = err
	weigh := ""

	for i := 0; i < len(arrayFileData)-4; i++ {
		for j := i + 1; j < len(arrayFileData)-3; j++ {
			for k := j + 1; k < len(arrayFileData)-2; k++ {
				for l := k + 1; l < len(arrayFileData)-1; l++ {
					for m := l + 1; m < len(arrayFileData); m++ {
						hand[0] = arrayFileData[i]
						hand[1] = arrayFileData[j]
						hand[2] = arrayFileData[k]
						hand[3] = arrayFileData[l]
						hand[4] = arrayFileData[m]

						if isStraightFlush(hand) {
							weigh = "Straight Flush"
							fmt.Println(hand, weigh)
						} else {
							if isFourOfAKind(hand) {
								weigh = "Four of a Kind"
								fmt.Println(hand, weigh)
							} else {
								if isFullHouse(hand) {
									weigh = "Full House"
									fmt.Println(hand, weigh)
								} else {
									if isFlush(hand) {
										weigh = "Flush"
										fmt.Println(hand, weigh)
									} else {
										if isStraight(hand) {
											weigh = "Straight"
											fmt.Println(hand, weigh)
										} else {
											if isThreeOfAKind(hand) {
												weigh = "Three of a Kind"
												fmt.Println(hand, weigh)
											} else {
												if isTwoPairs(hand) {
													weigh = "Two Pairs"
													fmt.Println(hand, weigh)
												} else {
													if isPair(hand) {
														weigh = "Pair"
														fmt.Println(hand, weigh)
													} else {
														fmt.Println(hand)
													}
												}
											}
										}
									}
								}
							}
						}
						if weigh != "" {
							resString := strings.Join(hand[:], ",") + " | " + weigh + "\n"
							resFile.WriteString(resString)
						}
						weigh = ""
					}
				}
			}
		}
	}
	_ = resFile.Close()
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

		go result(i)
	}

	time.Sleep(time.Second)
}
