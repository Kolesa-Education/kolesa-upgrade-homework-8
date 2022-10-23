package combinations

import (
	"strings"
)

type Counter map[byte]int
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

func getCategory(cards Hand) Category {
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
}
