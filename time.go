package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	o := "2019.8.12 ~ 2019.7.20"
	tmp := strings.Trim(strings.Split(o, "~")[1], " ")
	date, _ := time.Parse("2006.1.2", tmp)
	fmt.Println(tmp, date)
	ltime := "2019.08.04"
	// lt := ltime.Format(http.TimeFormat)
	ts, _ := time.Parse("20060102150405", time.Now().Format("20060102150405"))

	fmt.Println("now", time.Now())
	fmt.Println(ts)

	// fmt.Println(lt)
	tm2, _ := time.Parse("2006.01.02", ltime)
	ti := tm2.Format(http.TimeFormat)
	fmt.Println("time week", tm2.Weekday(), int(tm2.Weekday()))

	fmt.Println(ti)

	fmt.Println("getLastWednesday", getLastWednesday(date))
	fmt.Println("getLastWednesday", getLastWednesday(tm2))
	fmt.Println("getLastWednesday", getLastWednesday(time.Now()))
	trans := getLastWednesday(time.Now()).Format("2006.1.2")
	fmt.Println(trans)
	displayTimeEnd := getLastWednesday(time.Now())
	fmt.Println("displayTimeEnd", displayTimeEnd)
	displayTimeStart := getLastWednesday(displayTimeEnd).AddDate(0, 0, -7)
	fmt.Println("displayTimeStart", displayTimeStart)
	fmt.Println(displayTimeStart.Format("2006.1.2"), displayTimeEnd.Format("2006.1.2"))
	str := displayTimeStart.Format("2006.1.2") + " ~ " + displayTimeEnd.Format("2006.1.2")
	fmt.Println(str)

	formatedTime := getLastWednesday(time.Now())
	modifyTime := formatedTime.Format(http.TimeFormat)

	fmt.Println(modifyTime)
}

func getLastWednesday(date time.Time) (day time.Time) {
	var offset = [7]int{-4, -5, -6, 0, -1, -2, -3}
	week := int(date.Weekday())
	day = date.AddDate(0, 0, offset[week])
	return
}
