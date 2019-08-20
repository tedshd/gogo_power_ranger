package main

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type RecommendUser struct {
	AverageLikeCount int64    `json:"average_like_count"`
	RegCountry       string   `json:"reg_country"`
	UID              uint64   `json:"uid"`
	Nickname         string   `json:"nickname"`
	Avatar           string   `json:"avatar"`
	PostThumbnail    string   `json:"-"`
	LikeCount        uint     `json:"like_count" gorm:"column:belikeCount"`
	PostList         []string `json:"post_list"`
}

func main() {
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:Kingsoft@777@tcp(35.161.194.246:8000)/phototalk_adr")
	if err != nil {
		panic("failed to connect database")
	}

	var rl []RecommendUser
	country := ""
	err = db.Table("pg_index_user").Limit(250).Order("pg_userInfo.reg_country='" + country + "' desc").Order("pg_index_user.average_like_count desc").Select([]string{"DISTINCT pg_index_user.average_like_count", "pg_userInfo.reg_country", "pg_userInfo.uid", "pg_userInfo.nickname", "pg_userInfo.avatar", "pg_index_user.post_thumbnail", "pg_userInfo.belikeCount"}).Joins("left join pg_userInfo ON pg_index_user.uid = pg_userInfo.uid").Find(&rl).Error

	if err != nil {
		fmt.Println(err.Error())
	}
	for index, _ := range rl {
		var s = &rl[index]
		s.PostList = strings.Split(s.PostThumbnail, ",")
		fmt.Println(s)
	}
	// rl.PostList = strings.Split(rl.PostList, ",")
	// fmt.Println(rl)

	defer db.Close()

}
