package main

import (
	"fmt"
	"reflect"
)

type art struct {
	name string
	pid  int
}

type NewUser struct {
	Uid              int64   `json:"uid"`
	PostCount        int64   `json:"postCount"`
	FansCount        int64   `json:"fansCount"`
	RegCountry       string  `json:"reg_country"`
	CreateTime       int64   `json:"createTime"`
	BelikeCount      int64   `json:"belikeCount"`
	AverageLikeCount float32 `json:average_like_count`
	ext              *art
}

type d4 struct {
	name string `json:"name"`
	date int64 `json:date`
}

type Point struct {
    x float64
    y float64
}

func main() {
	arg := new(NewUser)
	arg.Uid = 123
	arg.PostCount = 33
	arg.FansCount = 22
	arg.RegCountry = "TW"
	update := new(art)
	update.name = "Ted"
	update.pid = 123456
	arg.ext = update

	run(arg)

	fmt.Println("===")

	fmt.Println(n())

	fmt.Println("===")

	d4Struct := d4{}
	fmt.Println(d4Struct)

	fmt.Println(reflect.TypeOf(d4Struct).Kind())

	fmt.Println("===")

	d4Array := []d4{}
	d4Array = append(d4Array, d4{
		name: "Ted",
		date: 1234})
	fmt.Println(d4Array)
	fmt.Println(reflect.TypeOf(d4Array).Kind())

	fmt.Println("===")

	p := Point{x: 3.0, y: 4.0}

    fmt.Println(p.x)
    fmt.Println(p.y)
}

func run(arg *NewUser) {
	fmt.Println("newUserList uid: %s\n", arg.Uid)
	fmt.Println("newUserList RegCountry: %s\n", arg.RegCountry)
	fmt.Println("newUserList AverageLikeCount: %s\n", arg.AverageLikeCount)
	fmt.Println("newUserList ext: %s\n", arg.ext)
}

func n() (res *NewUser) {
	p := new(NewUser)
	p.Uid = 4321
	res = p
	return
}
