package pipeline

import (
	"encoding/csv"
	"log"
	"os"
	"sync"

	"gonum.org/v1/gonum/stat/combin"
)

func readCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

func writeCsv(filePath string, data [][]string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("error creating a file: %s", err)
	}
	writer := csv.NewWriter(file)
	for _, oneComb := range data {
		err = writer.Write(oneComb)
		if err != nil {
			log.Fatalf("error writing to a file: %s", err)
		}
	}
	writer.Flush()
	file.Close()
}

func removeDuplicates(cards []string) []string {
	list := []string{}
	allKeys := make(map[string]bool)
	for _, item := range cards {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func getCombinations(cards []string, col int) [][]string {
	cs := combin.Combinations(len(cards), col)
	super_list := [][]string{}
	for _, c := range cs {
		list := []string{}
		for i := 0; i < col; i++ {
			list = append(list, cards[c[i]])
		}
		super_list = append(super_list, list)
	}
	return super_list
}

func Pipeline(num string, wg *sync.WaitGroup) {
	records := readCsv("dataset/dat" + num + ".csv")[0]
	records = removeDuplicates(records)
	combin_result := getCombinations(records, 5)
	//combin_result = combinations.getCategory(combin_result)
	writeCsv("results/data"+num+".csv", combin_result)
	wg.Done()
}
