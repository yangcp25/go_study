package main

import "fmt"

func main() {
	initData()
}

type Post struct {
	Id      int
	Title   string
	Content string
	Author  string
}

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}
func initData() {
	// 初始化
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{
		Id:      1,
		Title:   "x",
		Content: "x",
		Author:  "ycp",
	}

	post2 := Post{
		Id:      2,
		Title:   "xx",
		Content: "xx",
		Author:  "ycp1",
	}

	post3 := Post{
		Id:      3,
		Title:   "xxx",
		Content: "xxx",
		Author:  "ycp1",
	}

	post4 := Post{
		Id:      4,
		Title:   "xxxx",
		Content: "xxxx",
		Author:  "ycp2",
	}

	store(post1)
	store(post2)
	store(post4)
	store(post3)

	fmt.Println(PostById[1])
	fmt.Println(PostById[3])

	for author, post := range PostByAuthor["ycp1"] {
		fmt.Println(author, ":", post.Content)
	}

	PostById[1].Author = "ycp4"
	fmt.Println(PostById[1])
	delete(PostById, 1)

	fmt.Println(PostById[1])
	delete(PostByAuthor, "ycp")
	fmt.Println(PostById[1])
}
