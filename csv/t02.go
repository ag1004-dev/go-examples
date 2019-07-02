package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	myfilename := "t02.csv"
	processCsv(myfilename)
}

func getUnixTime(mytime string) string {
	const shortForm = "2006-01-02"
	t, _ := time.Parse(shortForm, mytime)
	//fmt.Println(t)
	//fmt.Println(t.Unix())
	data := strconv.FormatInt(t.Unix(), 10)
	return data
}

func getDate(date string) string {
	values := strings.Split(date, " ")
	return values[0]
}

func processRecord(record []string) {
	var output [3]string
	//fmt.Println(record)
	//fmt.Printf("%T\n", record)
	//fmt.Println(record[0])
	if strings.Contains(record[0], "datetime") == false {
		datestr := getDate(record[0])
		output[0] = getUnixTime(datestr)
	}
	if strings.Contains(record[4], "close") == false {
		output[1] = record[4]
	}
	if strings.Contains(record[0], "datetime") == false {
		fmt.Println(output[1], output[0])
	}
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
