package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//initData()
	//initJsonData()
	//initJsonData2()
	//initCsvData()
	initGobData()
}

func initGobData() {
	article := Article{
		Id:      1,
		Title:   "x",
		Summary: "x",
		Author:  "ycp",
	}

	writeGob(article, "article.gob")

	var articleData Article
	//读取
	readGob(&articleData, "article.gob")

	fmt.Println(articleData)
}

func readGob(a *Article, s string) {
	file, err := ioutil.ReadFile(s)

	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(file)
	encoder := gob.NewDecoder(buffer)
	err = encoder.Decode(a)

	if err != nil {
		panic(err)
	}
}

// 二进制写入文件
func writeGob(article Article, s string) {
	buffer := new(bytes.Buffer)
	encode := gob.NewEncoder(buffer)

	err := encode.Encode(article)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(s, buffer.Bytes(), 0600)

	if err != nil {
		panic(err)
	}
}

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
}

type Tutorial struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"summary"`
	Author  string `json:"author"`
}

func initCsvData() {
	csvFile, err := os.Create("tutorial.csv")

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	post1 := Tutorial{
		Id:      1,
		Title:   "x",
		Content: "x",
		Author:  "ycp",
	}

	post2 := Tutorial{
		Id:      2,
		Title:   "xx",
		Content: "xx",
		Author:  "ycp1",
	}

	post3 := Tutorial{
		Id:      3,
		Title:   "xxx",
		Content: "xxx",
		Author:  "ycp1",
	}

	post4 := Tutorial{
		Id:      4,
		Title:   "xxxx",
		Content: "xxxx",
		Author:  "ycp2",
	}

	tutorials := []Tutorial{
		post1,
		post2,
		post3,
		post4,
	}

	// 写入 UTF-8 BOM，防止中文乱码
	csvFile.WriteString("\xEF\xBB\xBF")

	wirter := csv.NewWriter(csvFile)

	for _, tutorial := range tutorials {
		line := []string{
			strconv.Itoa(tutorial.Id),
			tutorial.Title,
			tutorial.Content,
			tutorial.Author,
		}

		err := wirter.Write(line)

		if err != nil {
			panic(err)
		}
	}

	wirter.Flush()

	// 打开csv

	file, err := os.Open("tutorial.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	reader.FieldsPerRecord = -1

	record, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	var str2 []Tutorial

	for _, strings := range record {
		id, _ := strconv.ParseInt(strings[0], 0, 0)
		tutorial := Tutorial{
			int(id),
			strings[1],
			strings[2],
			strings[3],
		}
		str2 = append(str2, tutorial)
	}

	// 验证

	fmt.Println(str2[1].Id)
	fmt.Println(str2[1].Author)
	fmt.Println(str2[3].Id)
	fmt.Println(str2[3].Author)
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
