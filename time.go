package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	o := "2019.8.12 ~ 2019.8.20"
	tmp := strings.Trim(strings.Split(o, "~")[1], " ")
	date, _ := time.Parse("2016.1.2", tmp)
	fmt.Println(tmp)
	fmt.Println(date)
	ltime := "2019.08.02"
	// lt := ltime.Format(http.TimeFormat)
	ts, _ := time.Parse("20060102150405", time.Now().Format("20060102150405"))

	// t := time.Now()
	// ti := ts.Format(http.TimeFormat)
	fmt.Println(ts)

	// fmt.Println(lt)
	tm2, _ := time.Parse("2006.01.02", ltime)
	ti := tm2.Format(http.TimeFormat)
	// tm2, _ := time.Parse("01/02/2006", "02/08/2015")

	fmt.Println(ti)
}
