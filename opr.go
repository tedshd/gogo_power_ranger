package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	str := "thi si is"
	fmt.Println(str)
	af := strings.Trim(str, " ")
	fmt.Println(af)
	s := strings.Split("", ",")
	fmt.Println(s)
	fmt.Println(len(s))
}
