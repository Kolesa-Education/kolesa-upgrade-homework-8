package pockerCombination

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"sort"
	strconv "strconv"
)

type Combination struct {
	Сards []card.Card
}

func (c Combination) IsPair() (string, bool) {
	set := make(map[string]int)
	for _, v := range c.Сards {
		set[v.Face]++
		if set[v.Face] == 2 {
			return "Pair", true
		}
	}
	return "", false
}

func (c Combination) IsTwoPairs() (string, bool) {
	set := make(map[string]int)
	cnt := 0
	for _, v := range c.Сards {
		set[v.Face]++
		if set[v.Face] == 2 {
			cnt++
		}
	}
	if cnt >= 2 {
		return "Two Pairs", true
	}
	return "", false
}

func (c Combination) IsThreeOfAKind() (string, bool) {
	set := make(map[string]int)
	for _, v := range c.Сards {
		set[v.Face]++
		if set[v.Face] == 3 {
			return "Three of a Kind", true
		}
	}
	return "", false
}

func (c Combination) IsStraight() (string, bool) {
	sort.SliceStable(c.Сards, func(i, j int) bool {
		x, _ := strconv.Atoi(c.Сards[i].Face)
		y, _ := strconv.Atoi(c.Сards[j].Face)
		return x < y
	})
	cnt := 1
	n := len(c.Сards)
	for i := 0; i < n-1; i++ {
		cur, _ := strconv.Atoi(c.Сards[i].Face)
		nxt, _ := strconv.Atoi(c.Сards[i+1].Face)
		if cur+1 == nxt {
			cnt++
		}
	}
	if c.Сards[n-1].Face == "14" {
		cnt++
	}
	if cnt == 5 {
		return "Straight", true
	}
	return "", false
}

func (c Combination) IsFlush() (string, bool) {
	set := make(map[string]int)
	for _, v := range c.Сards {
		set[v.Suit]++
	}
	if len(set) == 1 {
		return "Flush", true
	}
	return "", false
}

func (c Combination) IsFullHouse() (string, bool) {
	set := make(map[string]int)
	three := false
	pair := false
	for _, v := range c.Сards {
		set[v.Face]++
		if set[v.Face] == 3 {
			three = true
		}
		if set[v.Face] == 2 {
			pair = true
		}
	}
	if three && pair {
		return "Full House", true
	}
	return "", false
}

func (c Combination) IsFourOfKind() (string, bool) {
	set := make(map[string]int)
	for _, v := range c.Сards {
		set[v.Face]++
		if set[v.Face] == 4 {
			return "Four of a kind", true
		}
	}
	return "", false
}

func (c Combination) IsStraightFlush() (string, bool) {
	title1, ok1 := c.IsFlush()
	title2, ok2 := c.IsStraight()
	if ok1 && ok2 {
		return title1 + " " + title2, true
	}
	return "", false
}

func (c Combination) GetTitle() (title string, ok bool) {
	if res, yes := c.IsPair(); yes {
		ok = true
		title = res
	}
	if res, yes := c.IsTwoPairs(); yes {
		ok = true
		title = res
	}

	if res, yes := c.IsThreeOfAKind(); yes {
		ok = true
		title = res
	}

	if res, yes := c.IsStraight(); yes {
		ok = true
		title = res
	}

	if res, yes := c.IsFlush(); yes {
		ok = true
		title = res
	}

	if res, yes := c.IsFullHouse(); yes {
		ok = true
		title = res
	}

	if res, yes := c.IsFourOfKind(); yes {
		ok = true
		title = res
	}

	if res, yes := c.IsStraightFlush(); yes {
		ok = true
		title = res
	}
	return title, ok
}
