package main

import "database/sql"

func main() {
	initSql()
}

var Db *sql.DB

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
}

func initSql() {
	var err error
	Db, err = sql.Open("mysql", "root:root@test_db?charset=uft8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
}

// 增
func (post *Post) Create() (err error) {
	sql := "insert into posts(title,content.author) values(?, ?, ?)"
	smt, err := Db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer smt.Close()

	res, err := Db.Exec(sql, post.Title, post.Summary, post.Author)

	if err != nil {
		panic(err)
	}

	postId, _ := res.LastInsertId()
	post.Id = int(postId)

	return
}

func getPost(id int) (post Post, err error) {
	post = Post{}

	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Summary, &post.Author)
	return
}
