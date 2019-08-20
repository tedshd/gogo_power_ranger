package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"phototalk:32cd5a3460sd564awdgd07@tcp(phototalk-adr-read.cq7ehxlwlj3o.us-west-2.rds.amazonaws.com:3306)/phototalk_adr")
	if err != nil {
		fmt.Printf("connect")
	}

	rows, err := db.Query("SELECT thrumbNail FROM pg_postInfo WHERE UID=560641164152537088 AND isPrivate=0 AND enable=1 ORDER BY createTime DESC limit ?", 3)
	if err != nil {
		log.Fatal(err)
	}

	//切記用完都要做 Close
	defer rows.Close()
	var arr []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s\n", strings.Replace(name, "%", "small", -1))
		arr = append(arr, strings.Replace(name, "%", "small", -1))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Print(strings.Join(arr, ","))
}
