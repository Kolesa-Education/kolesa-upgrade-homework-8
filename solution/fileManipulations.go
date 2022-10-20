package solution

import (
	"encoding/csv"
	"log"
	"os"
)

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()
	return file
}

func getDataFromCSV(fileName string) []string {
	file := readFile(fileName)
	csvReader := csv.NewReader(file)
	data, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}
	return data

}