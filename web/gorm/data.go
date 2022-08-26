package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func main() {
	initSql()
	initSql()
}

var Db *gorm.DB

func initSql() {
	var err error
	Db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})

	post := &Post{
		Title:   "x",
		Summary: "y",
		Author:  "yccp",
	}

	Db.Create(post)

	comment := Comment{
		Content: "嘻嘻",
		Author:  "wang",
	}
	Db.Model(&post).Association("Comments").Append(comment)

	//查询
	var postRes Post
	Db.Where("Author = ?", "yccp").First(&postRes)
	fmt.Println(postRes)

	//关联查询评论
	var commentsRes []Comment
	Db.Model(&postRes).Related(&commentsRes)
	fmt.Println(commentsRes)

}

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Author    string `json:"author"`
	CreatedAt time.Time
	Comments  []Comment
}

type Comment struct {
	Id        int
	Content   string
	Author    string
	PostId    int
	CreatedAt time.Time
}
