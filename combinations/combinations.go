package combinations

import (
	"sort"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
)

func IsPair(cards []card.Card) bool {

	prevFace := ""
	for i, firCard := range cards {
		counter := 0
		for _, secCard := range cards[i+1:] {
			if firCard.Face == secCard.Face {
				counter++

			}
		}
		if counter > 0 && prevFace != firCard.Face {
			if counter == 1 {
				return true
			} else {
				prevFace = firCard.Face

			}
		}
	}

	return false
}

func IsTwoPairs(cards []card.Card) bool {

	pairCounter := 0
	prevFace := ""
	for i, firCard := range cards {
		counter := 0
		for _, secCard := range cards[i+1:] {
			if firCard.Face == secCard.Face {
				counter++
			}
		}

		if counter > 0 && prevFace != firCard.Face {
			if counter == 1 {
				pairCounter++
			} else {
				prevFace = firCard.Face
			}
		}

	}
	if pairCounter == 2 {
		return true
	}

	return false
}

func IsThreeOfaKind(cards []card.Card) bool {

	prevFace := ""
	for i, firCard := range cards {
		counter := 0
		for _, secCard := range cards[i+1:] {
			if firCard.Face == secCard.Face {
				counter++

			}
		}
		if counter > 1 && prevFace != firCard.Face {
			if counter == 2 {
				return true
			} else {
				prevFace = firCard.Face

			}
		}
	}

	return false

}

func IsStraight(cards []card.Card) bool {

	sorted := make([]card.Card, 5)
	copy(sorted, cards)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Power < sorted[j].Power
	})

	// не разобрал круговую последовательность , когда ♣2,♣3,♣A,♣4,♣5 (больший элемент ссылается на меньший и так по кругу)

	for i := 0; i < len(sorted)-1; i++ {

		if sorted[i].Power+1 != sorted[i+1].Power {
			return false
		}

	}
	return true

}

func IsFlush(cards []card.Card) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Suit != cards[i+1].Suit {
			return false
		}
	}
	return true

}

func IsFullHouse(cards []card.Card) bool {

	if IsPair(cards) && IsThreeOfaKind(cards) {
		return true
	}
	return false

}
func IsFourOfaKind(cards []card.Card) bool {
	prevFace := ""
	for i, firCard := range cards {
		counter := 0
		for _, secCard := range cards[i+1:] {
			if firCard.Face == secCard.Face {
				counter++

			}
		}
		if counter > 2 && prevFace != firCard.Face {
			if counter == 3 {
				return true
			} else {
				prevFace = firCard.Face

			}
		}
	}

	return false
}
func IsStraightFlush(cards []card.Card) bool {
	if IsStraight(cards) && IsFlush(cards) {
		return true
	}
	return false
}
