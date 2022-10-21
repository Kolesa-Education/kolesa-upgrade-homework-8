package fileWorker

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
	"os"
)

func ReadCsv(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.Read()

	return records
}

func WriteCsv(path string, data []string) {
	var directory = strings.SplitAfterN(path, "/", 2)
	makeDirectoryIfNotExists(directory[0])

    file, err := os.Create(path)

		if err != nil {
			log.Fatalln("failed to open file", err)
		}

		writer := csv.NewWriter(file)
		if err = writer.Write(data); err != nil {
			log.Fatalln("error writing to a file!")
		}

		writer.Flush()
		_ = file.Close()
}

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	
	return nil
}
