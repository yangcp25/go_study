package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		return
	}

	res, _ := Db.Exec("insert into comments(content,author,post_id) values(?,?,?)", comment.Content, comment.Author, comment.Post.Id)

	id, _ := res.LastInsertId()
	comment.Id = int(id)
	return
}
