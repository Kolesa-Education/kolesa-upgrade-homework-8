package services

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/fileWorker"
)

const (
	InputDirectoryPath  = "dataset/"
	OutputDirectoryPath = "results/"
	FileFormat          = ".csv"
)

func Handle(fileName string) {
	data := fileWorker.ReadCsv(InputDirectoryPath + fileName + FileFormat)

	outputFile := OutputDirectoryPath + fileName + FileFormat
	fileWorker.WriteCsv(outputFile, data)
}
