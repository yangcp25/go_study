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

	setStruct(&s)
	fmt.Println(s.name1, s.Name2)

	s2 := test1{}
	setDefaultStruct(&s2)

}

type TestF struct {
	Val1 string
	Val2 string
}

func setDefaultStruct(t any) {
	val := reflect.ValueOf(t)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		fmt.Println("not struct2")
		return
	}

	val = val.Elem()
	tCate := val.Type()

	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		ts := tCate.Field(i)
		zero := reflect.Zero(ts.Type).Interface()

		if reflect.DeepEqual(v.Interface(), zero) {
			if v.CanSet() {
				var temp reflect.Value
				temp = reflect.ValueOf("ss")
				v.Set(temp)
			}
		}
	}

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

func setStruct(v any) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		fmt.Println("not struct2")
		return
	}

	fmt.Println("---------")

	val = val.Elem() // struct Value，可设置字段
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
		value2 := t.Field(i)
		if value.CanSet() {
			//value.SetString("tets")
			value.SetString("testxxx")
		}
		fmt.Println(value2.Type.String())
	}
}
