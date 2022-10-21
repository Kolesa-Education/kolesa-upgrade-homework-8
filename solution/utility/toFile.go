package utility

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"log"
	"os"
)

func ToFile(cards [5]card.Card, handType string, fileName string) {
	str := ""
	for i, card := range cards {
		str += card.Suit + card.Face
		if i == 4 {
			continue
		}
		str += ", "
	}

	str += " |" + handType

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(str + "\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
