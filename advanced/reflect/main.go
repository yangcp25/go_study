package main

import (
	"fmt"
	"reflect"
)

func main() {
	describe(123)
	describe("hello")
	describe([]int{1, 2, 3})
	describe(map[string]float64{"pi": 3.14})

	test := 123
	trySet(&test)

	fmt.Println(test)

	s := test1{
		name1: "test1",
		Name2: "test2",
	}

	getStruct(s)
}

func describe(v any) {
	val := reflect.ValueOf(v)

	fmt.Println(val.Type(), val.Kind(), val.Type().String(), val.Kind().String())
}

func trySet(data any) {
	val := reflect.ValueOf(data)
	v := val.Elem()
	v.SetInt(10)
}

type test1 struct {
	name1 string
	Name2 string `json:"name_1"`
}

func getStruct(v any) {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Struct {
		fmt.Println("not struct")
		return
	}

	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
		value2 := t.Field(i)
		if value.CanInterface() {
			fmt.Println(value2.Type, value.Interface())
		} else {
			fmt.Println(value.Type())
		}
	}
}
