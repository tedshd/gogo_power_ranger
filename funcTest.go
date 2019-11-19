package main

import "fmt"

func run(x string, n int) (s string) {
	if (n == 4) {
		n = 123
	}
	fmt.Println(x)
	fmt.Println(n)

	s = "RUN"

	return
}

func main() {
	run("str", 4)
	run("str", 234)
	// fmt.Println()
}