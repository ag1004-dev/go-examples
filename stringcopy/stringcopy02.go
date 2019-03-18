package main

import (
	"fmt"
	"log"
	"regexp"
)

func filterString(url string) (t bool) {
	matched, err := regexp.MatchString(`Hello`, url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(matched)
	return true
}

func copystring(a string) string {
	return (a + " ")[:len(a)]
}

func main() {
	a := "Hello, World!"
	b := copystring(a)
	result := filterString(a)
	if result {
		fmt.Println(b)
	}
}
