package main

import (
	"fmt"
	"reflect"
)

func main() {
	//v := &Test{}
	//DynamicCallFunc(v, "Add", 1, 2)
	//DynamicCallFunc(v, "Print", "hello")

	// 使用反射获取和修改函数小写字段

	testS := make([]int, 0)

	appendToSlice(&testS, 1)

	fmt.Println(testS)

	test := make(map[string]int)
	setMapKey(test, "test", 1)
	fmt.Println(test)

}

// 动态修改map slice 的值
func DynamicSetSlice(v any, value any) {
	val := reflect.ValueOf(v)
	if val.Elem().Kind() != reflect.Slice || val.Kind() != reflect.Ptr {
		panic("v must be a pointer")
	}
	val = val.Elem()
	newV := reflect.ValueOf(value)
	newSlice := reflect.Append(val, newV)
	val.Set(newSlice)
}

func DynamicSetMap(v any, key, val any) {

}

type Test struct{}

func (t Test) Add(a, b int) int {
	fmt.Println("Add", a, b)
	return a + b
}

func (t *Test) Print(a string) {
	fmt.Println("Print", a)
}

func DynamicCallFunc(v any, methodName string, args ...any) {
	val := reflect.ValueOf(v)
	method := val.MethodByName(methodName)
	if !method.IsValid() {
		fmt.Println("not method")
		return
	}
	//build params
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}
	method.Call(in)
}

func appendToSlice(v any, elm any) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		panic("v must be a slice")
	}
	val = val.Elem()
	newV := reflect.ValueOf(elm)
	val.Set(reflect.Append(val, newV))

}

func setMapKey(v any, key, value any) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Map {
		panic("v must be a map")
	}
	//val = val.Elem()
	newV := reflect.ValueOf(value)
	val.SetMapIndex(reflect.ValueOf(key), newV)
}

type test1 struct {
	Val int `json:"val"`
}
