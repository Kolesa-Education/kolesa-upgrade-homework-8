package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readDatasets(filePath string) ([]string, error) {
	// result := make(map[string]rune)
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(f)
	records, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func main() {
	temp, err := readDatasets("../dataset/dat0.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(temp)
}
