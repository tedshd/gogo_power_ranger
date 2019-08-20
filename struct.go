package main

import "fmt"

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
	fmt.Println(n())
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
