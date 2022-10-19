package main

import (
	"encoding/csv"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"os"
	"log"
	"io"
	"strings"
)




func readFile(fileName string){
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return file
}

func getDataFromCSV(fileName string){
	file = readFile(fileName)
	csvReader = csv.NewReader(file)
	data, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}
	return data[0]

}

func getCardMapFromSlice(dataSlice){
	dataMap = make(map[string]int, len(dataSlice))
	for i:=0; i<len(data_slice);i++ {
		if val, ok := dataMap[dataSlice[i]]; ok {
			dataSlice[dataSlice[i]]++
		}else{
			dataSlice[dataSlice[i]] = 0
		}
	}
	return dataMap
}

func checkQuantativeCombinations(dataSlice, dataMap){
	quantativeCombinations := make(map[string]string{
		"22": "2 pairs",
		"23": "full house",
		"32": "full house",
		"2": "pair",
		"3": "three of a kind",
		"4": "four of a kind"
	})
	combination := ""
	for i:=0;i<len(dataSlice);i++{
		switch dataMap[dataSlice[i]]{
		case 2:
			res += "2"
		case 3:
			res += "3"
		case 4:
			res += "4"
		}
	}
	res, ok = quantativeCombinations[combination]
	if ok{
		return res
	}
	return "0"

}

func checkStraightOrFlush(dataSlice){
	combination := ""
	cardsOrder = "2345678910jqka2345678910jqka"
	values [len(dataSlice)]string
	valuesString string
	suits := make(map[string]bool)
	for i:=0;i<len(dataSlice);i++ {
		suits[dataSlice[i][0]] = true
		values[i] = dataSlice[i][1]
	}
	if strings.Contains(cardsOrder, strings.Join(values, "")){
		combination += "straight"
	}

	if len(suits) == 1{
		combination += "flush"
	}
	return combination
}

func main() {
		fileName := ""
		data = getDataFromCSV(fileName)
		dataSlice = strings.Split(data, ',')
		dataMap = getCardMapFromSlice(dataSlice)

	}
}
