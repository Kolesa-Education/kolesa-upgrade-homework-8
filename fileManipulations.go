package main

import (
	"encoding/csv"
	"fmt"
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

func writeDataInCSV(combination []string, index int) {
	file, err := os.Create(fmt.Sprintf("results/dat%d.csv", index))

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	writer := csv.NewWriter(file)

	for _, data := range combination {
		row := []string{data + "\n"}
		_ = writer.Write(row)
	}

	/*if err = writer.Write(combination); err != nil {
		log.Fatalln("error writing to a file!")
	}*/

	writer.Flush()
	_ = file.Close()
}
