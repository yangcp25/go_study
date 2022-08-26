package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//initSql()
	initSql2()
}

func initSql2() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:13306)/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}

	// 测试
	post1 := &Post{
		Title:   "x",
		Summary: "y",
		Author:  "yccp",
	}
	comments := &Comment{
		Content: "x2",
		Author:  "ycp2",
		Post:    post1,
	}
	// 增加
	err = post1.Create()
	err = comments.Create()
	if err != nil {
		panic(err)
	}

	// 查询一个
	test, _ := getPost(post1.Id)

	fmt.Println(test)

}

var Db *sql.DB

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Author   string `json:"author"`
	Comments []Comment
}

func initSql() {
	var err error
	Db, err = sql.Open("mysql", "root:root@/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}

	// 测试
	post1 := &Post{
		Title:   "x",
		Summary: "y",
		Author:  "yccp",
	}
	post2 := &Post{
		Title:   "x2",
		Summary: "y2",
		Author:  "ycp2",
	}
	// 增加
	err = post1.Create()
	err = post2.Create()
	if err != nil {
		panic(err)
	}

	// 查询
	posts, _ := getPosts(10)

	for _, post := range posts {
		fmt.Println(post.Id, post.Title, post.Author)
	}

	// 修改
	post2.Summary = "yyyy"
	err = post2.updatePosts()
	if err != nil {
		panic(err)
	}

	// 查询一个
	test, _ := getPost(post2.Id)

	fmt.Println(test)

	// 删除

	post1.deletePosts()

	// 查询
	posts, _ = getPosts(10)

	for _, post := range posts {
		fmt.Println(post.Id, post.Title, post.Author)
	}
}

// Create 增
func (post *Post) Create() (err error) {
	sql := "insert into posts(title,content,author) values(?, ?, ?)"
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
	//关联表
	comments, _ := Db.Query("select content,author from comments where post_id=?", post.Id)

	for comments.Next() {
		comment := Comment{}
		err := comments.Scan(&comment.Content, &comment.Author)
		if err != nil {
			panic(err)
		}
		post.Comments = append(post.Comments, comment)
	}
	return
}

func getPost2(id int) (post Post, err error) {
	row, err := Db.Query("select id,content,author from posts  where id = ? limit 1", id)
	if err != nil {
		panic(err)
	}
	post = Post{}

	for row.Next() {
		post := &Post{}
		err = row.Scan(&post.Id, &post.Summary, &post.Author)
		if err != nil {
			panic(err)
		}
	}
	return
}

// 获取数组
func getPosts(limit int) (results []*Post, err error) {
	pre, err := Db.Prepare("select id,content,author from posts limit ?")
	if err != nil {
		panic(err)
	}

	defer pre.Close()

	// result 是一个迭代器
	result, err := pre.Query(limit)

	if err != nil {
		panic(err)
	}

	for result.Next() {
		post := &Post{}
		err = result.Scan(&post.Id, &post.Summary, &post.Author)
		if err != nil {
			panic(err)
		}
		results = append(results, post)
	}
	return
}

// 修改
func (post *Post) updatePosts() (err error) {
	row, err := Db.Prepare("update posts set content=?,author=? where id = ?")
	if err != nil {
		panic(err)
	}
	_, err = row.Exec(post.Summary, post.Author, post.Id)
	if err != nil {
		return err
	}
	return
}

// 删除
func (post *Post) deletePosts() (err error) {
	row, err := Db.Prepare("delete from posts where id = ?")
	if err != nil {
		panic(err)
	}
	_, err = row.Exec(post.Id)
	if err != nil {
		return err
	}
	return
}
