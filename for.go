package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	filter := ""
	for index := 0; index < len(arr); index++ {
		filter = filter + "AND uid != " + strconv.Itoa(arr[index]) + " "
	}
	fmt.Println(filter)

	for index := 0; index < len(arr); index++ {
		run(arr[index])
	}

}

func run(n int) {
	if n == 3 {
		return
	}
	fmt.Println(n)
}
