package services

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/fileWorker"
	"sync"
)

const (
	InputDirectoryPath  = "dataset/"
	OutputDirectoryPath = "results/"
	FileFormat          = ".csv"
)

func Handle(fileName string, wg *sync.WaitGroup) {
	data := fileWorker.ReadCsv(InputDirectoryPath + fileName + FileFormat)

	outputFile := OutputDirectoryPath + fileName + FileFormat
	fileWorker.WriteCsv(outputFile, data)

	wg.Done()
}
