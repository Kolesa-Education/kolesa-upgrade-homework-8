package main

import (
	"encoding/csv"
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"os"
	"log"
	"io"
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
	return data

}

func main() {
		fileName := ""
		data = getDataFromCSV(fileName)
	}
}
