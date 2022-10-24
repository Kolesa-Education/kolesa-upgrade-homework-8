package main

import (
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/pipeline"
	"strconv"
	"sync"
	"time"
)

func createPokerCombinations() {
	startTime := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go pipeline.Pipeline(strconv.Itoa(i), &wg)
	}

	wg.Wait()
	elapsedTime := time.Since(startTime)
	fmt.Print("Затрачено времени на выполнение задания ", elapsedTime)
}
