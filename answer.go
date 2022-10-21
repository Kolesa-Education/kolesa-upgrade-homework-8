package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	// "io/ioutil"
)

func main() {
	// Read Filnames in folder
	filename := []string{}
	f, err := os.Open("dataset")

	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		filename = append(filename, v.Name())
		//fmt.Println(v.Name(), v.IsDir())
	}

	fmt.Println(filename)
	// Reading one file
	// Read all the files between 0 to 99
	// Better way to read

	// open file
	for _, fname := range filename {
		//for fnum := 0; fnum < 100; fnum++{
		f, err = os.Open("dataset/" + fname) //+strconv.Itoa(fnum)+".csv")
		if err != nil {
			log.Fatal(err)
		}

		//-------------------------------------
		// read csv values using csv.Reader
		csvReader := csv.NewReader(f)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		// convert records to array of structs
		//shoppingList := createShoppingList(data)

		// print the array
		fmt.Printf("%+v\n", data)
		//-------------------------------------------------------

		f.Close()

		//}
	}
	writeFile(filename)
}

func writeFile(filename []string) {
	for i := 0; i < len(filename); i++ {
		file, err := os.Create(fmt.Sprintf("results/result%d.csv", i))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = file.WriteString("hi")
		if err != nil {
			panic(err)
		}
	}
}
