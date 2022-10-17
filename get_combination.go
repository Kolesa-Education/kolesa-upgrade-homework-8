package main

import (
	"strconv"
	"sync"

	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/pipeline"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go pipeline.Pipeline(strconv.Itoa(i), &wg)
		wg.Wait()
	}
}
