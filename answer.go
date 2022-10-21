package main

import (
	"encoding/csv"
	"fmt"

	//"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	//"github.com/samber/lo"
	"log"
	//"strconv"
	//"math/rand"
	"os"
	// /"strings"
)

type CardList struct {
	Cards []Card
	Combo string
}

type Card struct {
	Suite string
	Face  string
}

// Read Function
func readCards(folderName string) []CardList {

	filename := []string{}
	var suite string
	var face string
	card := []Card{}
	cardlist := []CardList{}
	var oldFname = ""
	f, err := os.Open(folderName)

	if err != nil {
		fmt.Println(err)

	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)

	}

	for _, v := range files {
		filename = append(filename, v.Name())
	}

	// open file
	for i, fname := range filename {

		if i == 0 {
			oldFname = fname

		}
		if oldFname != fname {
			oldFname = fname
			//fmt.Printf("------------------\n")
		}
		f, err = os.Open(folderName + "/" + fname) //+strconv.Itoa(fnum)+".csv")
		if err != nil {
			log.Fatal(err)
		}
		csvReader := csv.NewReader(f)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		dat := data[0]
		for _, d := range dat {
			if len(d) > 4 {
				suite = d[:len(d)-2]
				face = d[3:len(d)]
			} else {

				suite = d[:len(d)-1]
				face = d[3:len(d)]
			}
			card = append(card, Card{suite, face})

		}

		cardlist = append(cardlist, CardList{card, ""})
		card = []Card{}

		f.Close()

	}
	fmt.Printf("%+v\n", cardlist[0])
	return cardlist

}

func searchPoker(clist []CardList) {
	cSpade := 0
	cHeart := 0
	cClub := 0
	cDiamond := 0
	totLen := 0
	c2 := 0
	c3 := 0
	c4 := 0
	c5 := 0
	c6 := 0
	c7 := 0
	c8 := 0
	c9 := 0
	c10 := 0
	cJ := 0
	cQ := 0
	cK := 0
	cA := 0
	straightFlag := false
	combo := ""
	str := ""

	for _, cardl := range clist {
		fmt.Printf("------------------\n")

		totLen = len(cardl.Cards)
		for c_i, c := range cardl.Cards {
			if c.Suite == "\u2665" { // Heart
				cHeart = cHeart + 1

			} else if c.Suite == "\u2663" { // Clubs
				cClub = cClub + 1

			} else if c.Suite == "\u2666" { // Diamond
				cDiamond = cDiamond + 1

			} else if c.Suite == "\u2660" { // Spade
				cSpade = cSpade + 1
			}

			if c.Face == "A" {
				cA = cA + 1
			} else if c.Face == "2" {
				c2 = c2 + 1
			} else if c.Face == "3" {
				c3 = c3 + 1
			} else if c.Face == "4" {
				c4 = c4 + 1
			} else if c.Face == "5" {
				c5 = c5 + 1
			} else if c.Face == "6" {
				c6 = c6 + 1
			} else if c.Face == "7" {
				c7 = c7 + 1
			} else if c.Face == "8" {
				c8 = c8 + 1
			} else if c.Face == "9" {
				c9 = c9 + 1
			} else if c.Face == "10" {
				c10 = c10 + 1
			} else if c.Face == "J" {
				cJ = cJ + 1
			} else if c.Face == "Q" {
				cQ = cQ + 1
			} else if c.Face == "K" {
				cK = cK + 1
			}

			// Find Consecutive
			if totLen >= 5 && c_i < totLen-5 {
				if cardl.Cards[c_i].Suite == cardl.Cards[c_i+1].Suite && cardl.Cards[c_i].Suite == cardl.Cards[c_i+2].Suite && cardl.Cards[c_i].Suite == cardl.Cards[c_i+3].Suite && cardl.Cards[c_i].Suite == cardl.Cards[c_i+4].Suite {
					straightFlag = true
				}
			}

			str = str + "," + c.Suite + c.Face

		}

		// Straight Flush
		if (cSpade >= 5 || cDiamond >= 5 || cHeart >= 5 || cClub >= 5) && straightFlag {
			combo = "STRAIGHT FLUSH"

		} else if cA == 4 || c2 == 4 || c3 == 4 || c4 == 4 || c5 == 4 || c6 == 4 || c7 == 4 || c8 == 4 || c9 == 4 || c10 == 4 || cJ == 4 || cQ == 4 || cK == 4 {
			//Full House
			if cA == 2 || c2 == 2 || c3 == 2 || c4 == 2 || c5 == 2 || c6 == 2 || c7 == 2 || c8 == 2 || c9 == 2 || c10 == 2 || cJ == 2 || cQ == 2 || cK == 2 {
				combo = "FOUR OF A KIND"
			}
			combo = "FOUR OF A KIND"
		} else if (cSpade >= 5 || cDiamond >= 5 || cHeart >= 5 || cClub >= 5) && straightFlag {
			combo = "FLUSH"

		} else if straightFlag {
			combo = "STRAIGHT"

		} else if cA == 3 || c2 == 3 || c3 == 3 || c4 == 3 || c5 == 3 || c6 == 3 || c7 == 3 || c8 == 3 || c9 == 3 || c10 == 3 || cJ == 3 || cQ == 3 || cK == 3 {

			combo = "THREE OF A KIND"
		} else if cA == 2 || c2 == 2 || c3 == 2 || c4 == 2 || c5 == 2 || c6 == 2 || c7 == 2 || c8 == 2 || c9 == 2 || c10 == 2 || cJ == 2 || cQ == 2 || cK == 2 {
			if cA == 2 || c2 == 2 || c3 == 2 || c4 == 2 || c5 == 2 || c6 == 2 || c7 == 2 || c8 == 2 || c9 == 2 || c10 == 2 || cJ == 2 || cQ == 2 || cK == 2 {
				combo = "TWO PAIR"
			}
			combo = "PAIR"
		}

		fmt.Println(combo)

		straightFlag = false

		cSpade = 0
		cHeart = 0
		cClub = 0
		cDiamond = 0
		totLen = 0
		c2 = 0
		c3 = 0
		c4 = 0
		c5 = 0
		c6 = 0
		c7 = 0
		c8 = 0
		c9 = 0
		c10 = 0
		cJ = 0
		cQ = 0
		cK = 0
		cA = 0

	}

}

// Write Function
func main() {
	// Read Filnames in folder
	//clist := []CardList{}
	clist := readCards("dataset")
	searchPoker(clist)
	// writeFile(filename)
}

// func writeFile(filename []string) {
// 	for i := 0; i < len(filename); i++ {
// 		file, err := os.Create(fmt.Sprintf("results/result%d.csv", i))
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer file.Close()
// 		_, err = file.WriteString("hi")
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
