package main

import (
	"fmt"
)

func copys1(a string) string {
	return (a + " ")[:len(a)]
}

func copys2(a string) string {
	if len(a) == 0 {
		return ""
	}
	return a[0:1] + a[1:]
}

func main() {
	a := "Hello, World!"
	b := a[:5]
	c := copys1(b)
	d := copys2(b)
	fmt.Println(a, b, c, d)
}
