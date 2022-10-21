package combinations

import "github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"

const (
	Pair          = "Pair"
	TwoPair       = "Two Pairs"
	ThreeOfKind   = "Three of a Kind"
	Straight      = "Straight"
	Flush         = "Flush"
	FullHouse     = "Full House"
	FourOfKind    = "Four of a kind"
	StraightFlush = "Straight Flush"
)

func cardFaceToMap(cards []card.Card) map[string]int {
	m := make(map[string]int)

	for i := 0; i < 5; i++ {
		if _, ok := m[cards[i].Face]; ok {
			m[cards[i].Face] += 1
			continue
		}
		m[cards[i].Face] = 1
	}
	return m
}
func cardSuitToMap(cards []card.Card) map[string]int {
	m := make(map[string]int)

	for i := 0; i < 5; i++ {
		if _, ok := m[cards[i].Suit]; ok {
			m[cards[i].Suit] += 1
			continue
		}
		m[cards[i].Suit] = 1
	}
	return m
}

func isFlush(cards []card.Card) bool {
	for i := 1; i < 5; i++ {
		if cards[0].Suit != cards[i].Suit {
			return false
		}
	}
	return true
}
func isStraight(cards []card.Card) bool {
	for i := 0; i < 5; i++ {
		face := cards[i].Face
		check := false
		for j := 0; j < 5; j++ {
			switch face {
			case card.Face2:
				if containFace(cards, card.FaceAce) || containFace(cards, card.Face3) {
					check = true
				}
			case card.Face3:
				if containFace(cards, card.Face2) || containFace(cards, card.Face4) {
					check = true
				}
			case card.Face4:
				if containFace(cards, card.Face3) || containFace(cards, card.Face5) {
					check = true
				}
			case card.Face5:
				if containFace(cards, card.Face4) || containFace(cards, card.Face6) {
					check = true
				}
			case card.Face6:
				if containFace(cards, card.Face5) || containFace(cards, card.Face7) {
					check = true
				}
			case card.Face7:
				if containFace(cards, card.Face8) || containFace(cards, card.Face6) {
					check = true
				}
			case card.Face8:
				if containFace(cards, card.Face8) || containFace(cards, card.Face7) {
					check = true
				}
			case card.Face9:
				if containFace(cards, card.Face8) || containFace(cards, card.Face10) {
					check = true
				}
			case card.Face10:
				if containFace(cards, card.Face9) || containFace(cards, card.FaceJack) {
					check = true
				}
			case card.FaceJack:
				if containFace(cards, card.Face10) || containFace(cards, card.FaceQueen) {
					check = true
				}
			case card.FaceQueen:
				if containFace(cards, card.FaceJack) || containFace(cards, card.FaceKing) {
					check = true
				}
			case card.FaceKing:
				if containFace(cards, card.FaceAce) || containFace(cards, card.FaceQueen) {
					check = true
				}
			case card.FaceAce:
				if containFace(cards, card.Face2) || containFace(cards, card.FaceKing) {
					check = true
				}
			}
		}
		if !check {
			return false
		}

	}

	return true
}
func isStraightFlush(cards []card.Card) bool {
	if isFlush(cards) && isStraight(cards) {
		return true
	}
	return false
}

func GetCombination(cards []card.Card) string {
	mapOfFaces := cardFaceToMap(cards)
	mapOfSuits := cardSuitToMap(cards)
	if isStraightFlush(cards) {
		return StraightFlush
	}
	for _, val := range mapOfFaces {
		if val == 4 {
			return FourOfKind
		}
		if (val == 3 || val == 2) && len(mapOfFaces) == 2 {
			return FullHouse
		}
	}
	if len(mapOfSuits) == 1 {
		return Flush
	}
	if isFlush(cards) {
		return Flush
	}
	if isStraight(cards) {
		return Straight
	}
	for _, val := range mapOfFaces {
		if val == 3 {
			return ThreeOfKind
		}
	}
	count := 0
	for _, val := range mapOfFaces {
		if val == 2 {
			count++
		}
	}
	if count == 2 {
		return TwoPair
	} else if count == 1 {
		return Pair
	}
	return ""
}

func containFace(cards []card.Card, face string) bool {
	for i := 0; i < 5; i++ {
		if cards[i].Face == face {
			return true
		}
	}
	return false
}
