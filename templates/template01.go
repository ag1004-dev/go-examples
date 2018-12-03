package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	fmt.Println("Hello, playground")
	type Inventory struct {
		Material string
		Count    uint
	}
	sweater := Inventory{"wool", 17}
	bibs := Inventory{"cotton", 5}
	tmpl, err := template.New("testme").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweater)
	fmt.Println("\nand")
	err = tmpl.Execute(os.Stdout, bibs)
	fmt.Println("\nbye")
	if err != nil {
		panic(err)
	}
}
