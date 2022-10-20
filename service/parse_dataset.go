package service

import (
	"io/ioutil"
	"strings"
)

func parseCardsSlice(filepath string) ([]string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	tempSlice := string(data)
	cardsSlice := deleteDuplicate(strings.Split(tempSlice[:len(tempSlice)-1], ","))

	return cardsSlice, nil

}

func deleteDuplicate(cards []string) []string {
	tempMap := make(map[string]bool)
	result := []string{}
	for _, item := range cards {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = true
			result = append(result, item)
		}
	}
	return result
}
