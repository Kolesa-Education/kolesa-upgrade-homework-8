package combinations

import (
	"sort"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

/*type Counter map[byte]int
type Category byte
type Hand []string

const (
	HighCard      Category = iota // highest card
	OnePair                       // two cards of same value
	TwoPairs                      // two different pairs
	ThreeOfAKind                  // three cards of same value
	Straight                      // all cards are in consecutive values
	Flush                         // all cards of same suit
	FullHouse                     // three of a kind and a pair
	FourOfAKind                   // four cards of same value
	StraightFlush                 // all cards consecutive and same suit
)

const faces = "23456789TJQKA"

func CardRank(b byte) int {
	return strings.Index(faces, string(b))
}

func Len(cards Hand) int {
	return 5
}

func Swap(cards Hand, i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}

func Less(cards Hand, i, j int) bool {
	return CardRank(cards[i][1]) < CardRank(cards[j][1])
}

func isFlush(cards Hand) bool {
	firstSuite := cards[0][0]
	for _, card := range cards {
		if card[0] != firstSuite {
			return false
		}
	}
	return true
}

func isStraight(cards Hand) bool {
	var values []byte
	for _, card := range cards {
		values = append(values, card[1])
	}
	if strings.Contains(faces, string(values)) {
		return true
	}
	return false
}

func Has(c Counter, num int) bool {
	for _, count := range c {
		if count == num {
			return true
		}
	}
	return false
}

func getValue(c Counter, num int) (values []byte) {
	for val, count := range c {
		if count == num {
			values = append(values, val)
		}
	}
	return
}

func getCategory(cards []string) Category {
	//sort.Sort(cards)

	if isStraight(cards) {
		if isFlush(cards) {
			return StraightFlush
		}
		return Straight
	}
	if isFlush(cards) {
		return Flush
	}

	CardCounter := make(Counter)
	for _, card := range cards {
		CardCounter[card[1]]++
	}
	for _, n := range []int{4, 3, 2} {
		if Has(CardCounter, n) {
			switch n {
			case 4:
				return FourOfAKind
			case 3:
				if Has(CardCounter, 2) {
					return FullHouse
				} else {
					return ThreeOfAKind
				}
			case 2:
				if len(CardCounter) == 3 {
					pairs := getValue(CardCounter, 2)
					if pairs[0] > pairs[1] {
						pairs[0], pairs[1] = pairs[1], pairs[0]
					}
					return TwoPairs
				} else {
					return OnePair
				}
			}
		}
	}
	return HighCard
}*/

func GetStraightFlush(cards []card.Card) bool {
	straight := GetStraight(cards)
	if !straight {
		return false
	}
	flush := GetFlush(cards)
	if !flush {
		return false
	}
	return true
}

func GetFourOfAKind(cards []card.Card) bool {
	var counter int
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				counter++
			}
			if counter == 3 {
				return true
			}
		}
		counter = 0
	}
	return false
}

func GetFullHouse(cards []card.Card) bool {
	threeOfAKind := GetThreeOfAKind(cards)
	if !threeOfAKind {
		return false
	}
	pair := GetPair(cards)
	if !pair {
		return false
	}
	return true
}

func GetFlush(cards []card.Card) bool {
	firstCard := cards[0]
	for _, compCard := range cards[+1:] {
		if firstCard.Suit != compCard.Suit {
			return false
		}
	}
	return true
}

func GetStraight(cards []card.Card) bool {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Weight < cards[j].Weight
	})
	firstCard := cards[0]
	lastCard := cards[len(cards)-1]
	if lastCard.Weight == 14 {
		check := checkAces(cards)
		if !check {
			return false
		}
		lastCard.Weight = 1
		cardsWithoutAce := cards[:len(cards)-1]
		cardsAceFirst := make([]card.Card, 1)
		cardsAceFirst[0] = lastCard
		cardsAceFirst = append(cardsAceFirst, cardsWithoutAce...)
		aceFirst := GetStraight(cardsAceFirst)
		if aceFirst {
			return true
		}
	}
	for i, compCard := range cards[1:] {
		if compCard.Weight != firstCard.Weight+(i+1) {
			return false
		}
	}
	return true
}

func checkAces(cards []card.Card) bool {
	cpCards := make([]card.Card, 0)
	copy(cpCards, cards)
	sort.Slice(cpCards, func(i, j int) bool {
		return cards[i].Weight > cards[j].Weight
	})
	firstAce := cards[0]
	for _, curCard := range cards[1:] {
		if firstAce.Weight == curCard.Weight {
			return false
		}
	}
	return true
}

func GetThreeOfAKind(cards []card.Card) bool {
	var counter int
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				counter++
			}
			if counter == 2 {
				return true
			}
		}
		counter = 0
	}
	return false
}

func GetTwoPairs(cards []card.Card) bool {
	var firstPair bool
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				if !firstPair {
					firstPair = true
					continue
				}
				return true
			}
		}
	}
	return false
}

func GetPair(cards []card.Card) bool {
	for i, curCard := range cards {
		for _, compCard := range cards[i+1:] {
			if curCard.Face == compCard.Face {
				return true
			}
		}
	}
	return false
}
