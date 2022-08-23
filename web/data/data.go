package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//initData()
	//initJsonData()
	initJsonData2()
	//initCsvData()
	//initGobData()
}

func initJsonData2() {
	// 初始化
	var books map[int]*Book = make(map[int]*Book)
	book1 := &Book{Id: 1, Title: "Title", Summary: "Summary", Author: "ycp"}
	books[book1.Id] = book1

	//json 序列化
	data, _ := json.Marshal(books)

	file1, _ := os.Create("books.json")
	defer file1.Close()
	_, err := file1.Write(data)

	if err != nil {
		panic(err)
	}

	file2, _ := os.Open("books.json")
	defer file2.Close()
	dataEncode := make([]byte, len(data))
	file2.Read(dataEncode)

	var booksEncode map[int]*Book
	json.Unmarshal(dataEncode, &booksEncode)

	fmt.Printf("%#v", booksEncode[book1.Id])
	//fmt.Println(booksEncode)
}

func initCsvData() {

}
func initGobData() {

}

type Book struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
}

func initJsonData() {
	// 初始化
	var books map[int]*Book = make(map[int]*Book)
	book1 := &Book{Id: 1, Title: "Title", Summary: "Summary", Author: "ycp"}
	books[book1.Id] = book1

	//json 序列化
	data, _ := json.Marshal(books)
	err := ioutil.WriteFile("books.json", data, 0644)

	if err != nil {
		panic(err)
	}

	dataEncode, _ := ioutil.ReadFile("books.json")
	var booksEncode map[int]*Book
	json.Unmarshal(dataEncode, &booksEncode)

	fmt.Println(booksEncode)
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
