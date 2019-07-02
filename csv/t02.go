package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	myfilename := "t02.csv"
	processCsv(myfilename)
}

func processRecord(record []string) {
	fmt.Println(record)
	fmt.Printf("%T\n", record)
	fmt.Println(record[0])
}

func processCsv(filename string) {

	dat, err := ioutil.ReadFile(filename)
	check(err)
	token := string(dat)

	r := csv.NewReader(strings.NewReader(token))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		processRecord(record)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
