package main

import "fmt"

func main() {
	fmt.Println("Внимание, директории /dataset и /results должны быть созданы заранее")
	createDataset()
	createPokerCombinations()
}
