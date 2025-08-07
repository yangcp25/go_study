package main

import "fmt"

type Person struct {
	Total *int64 `json:"total"`
}

type Test struct {
	Person Person `json:"person"`
}

func (receiver *Person) Set(data *int64) {
	receiver.Total = data
}

func main() {
	//p := Person{}
	//t := Test{
	//	Person: p,
	//}
	//t.Person.Set(100)
	//fmt.Println(p.Total)
	//fmt.Println(t.Person.Total)
	test()
}

func test() {
	p := Person{}
	test2(p)
	fmt.Println(p.Total)
}

func test2(p Person) {
	val := int64(111)
	p.Set(&val)
}
