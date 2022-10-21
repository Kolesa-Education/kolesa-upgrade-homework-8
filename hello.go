package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"gonum.org/v1/gonum/stat/combin"
)

func Pair(allCombin string) bool {
	r := []rune(allCombin)

	group := []string{string(r[1]), string(r[3]), string(r[5]), string(r[7]), string(r[9])}
	for i1, v := range group {
		for i2 := i1 + 1; i2 < len(group); i2++ {
			if v == group[i2] {
				return true
			}
		}
	}
	return false
}
func writeCsvCombination() {
	file, _ := os.Open(fmt.Sprintf("dataset/dat0.csv"))
	r := csv.NewReader(file)
	record, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	justString := strings.Join(record, ",")
	le := strings.Split(justString, ",")
	cb := combin.Combinations(len(le), 5)
	for _, v := range cb {
		allCombin := fmt.Sprintf("%s%s%s%s%s", le[v[0]], le[v[1]], le[v[2]], le[v[3]], le[v[4]])
		if Pair(allCombin) {
			writeCsv(allCombin, "Pair")
		}
	}  
}

//var curId int

func writeCsv(allCombin string, combination_type string) {
	file2, err := os.Create(fmt.Sprintf("results/combinella1.csv"))
	if err != nil {
		fmt.Println(err.Error())
	}
	w := csv.NewWriter(file2)
	allCombin = allCombin + ";" + combination_type
	for _, records := range []string{allCombin} { 
		//fmt.Println(records)
		if err = w.Write([]string{records}); err != nil {
			log.Fatalln("error writing to a file!")
		}
		//fmt.Printf("%T",[]string{records} )
	}
	w.Flush()
	file2.Close()
	//curId = curId + 1
}

func main() {
	writeCsvCombination()
}
