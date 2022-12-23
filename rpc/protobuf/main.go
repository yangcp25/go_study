package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	_ "protoc_demo/proto"
	articleService "protoc_demo/proto"
)

func main() {
	articles := articleService.Article{
		User:    "ycp",
		Id:      1,
		Readers: []string{"1", "读者2"},
	}
	//fmt.Printf("%v", articles)
	fmt.Printf("%v", articles.GetReaders())
	data, _ := proto.Marshal(&articles)
	fmt.Println(data)

	test := articleService.Article{}

	proto.Unmarshal(data, &test)

	fmt.Println(test)
}
