package main

import "fmt"

type DATA struct {
	name string
	id   int
}

func main() {
	var res interface{}
	res = run()
	fmt.Println(res)
	handle(res)

	fmt.Println()
}

func run() (data *DATA) {
	tmp := new(DATA)
	tmp.id = 321
	tmp.name = "TT"
	data = tmp
	return
}

func handle(arg interface{}) {
	tmp := arg.(*DATA)
	fmt.Println(tmp.id)
}
