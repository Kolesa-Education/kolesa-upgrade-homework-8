package combinations

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/solution/utility"
)

func Combinations(inputArr []card.Card, outputArr [5]card.Card,
	start, end, index, r int, outputFile string) {
	if index == r {
		Process(outputArr, outputFile)
		return
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
		outputArr[index] = inputArr[i]
		Combinations(inputArr, outputArr, i+1, end, index+1, r, outputFile)
	}
}

func Process(cards [5]card.Card, filename string) {
	if IsStraightFlush(cards) {
		utility.ToFile(cards, "straight flush", filename)
	} else if IsFourKind(cards) {
		utility.ToFile(cards, "four kind", filename)
	} else if IsFullHouse(cards) {
		utility.ToFile(cards, "full house", filename)
	} else if IsFlush(cards) {
		utility.ToFile(cards, "flush", filename)
	} else if IsStraight(cards) {
		utility.ToFile(cards, "straight", filename)
	} else if IsThreeKind(cards) {
		utility.ToFile(cards, "three kind", filename)
	} else if IsTwoPairs(cards) {
		utility.ToFile(cards, "two pairs", filename)
	} else if IsPair(cards) {
		utility.ToFile(cards, "pair", filename)
	} else {
		return
	}
}
