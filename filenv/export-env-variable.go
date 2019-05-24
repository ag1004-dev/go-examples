package main

import "fmt"
import "io/ioutil"
import "os"

func main() {

	dat, err := ioutil.ReadFile("/Users/ma/.influxdbv2/credentials")
	check(err)

	token := string(dat)
	part1 := "export INFLUX_TOKEN="

	result := fmt.Sprintf("%s%s%s", part1, token, "\n")
	// fmt.Println(result)

	f, err := os.Create("/Users/ma/.influxenv")
	check(err)
	defer f.Close()

	n3, err := f.WriteString(result)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
