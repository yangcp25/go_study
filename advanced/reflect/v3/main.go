package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 动态创建和调用函数
	// 定义一个函数类型：func(int, int) int
	fnType := reflect.FuncOf([]reflect.Type{reflect.TypeOf(3), reflect.TypeOf(5)}, []reflect.Type{reflect.TypeOf(0)}, false)
	// 使用 MakeFunc 创建一个函数：计算两个整数之和
	sumVal := reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
		a := args[0].Int()
		b := args[1].Int()
		// 将 int64 转换为 int
		sum := int(a + b)
		return []reflect.Value{reflect.ValueOf(sum)}
	})
	// 将 reflect.Value 转回接口
	sumFunc := sumVal.Interface().(func(int, int) int)
	fmt.Println("3+5=", sumFunc(3, 5))

	// 模仿gorm 生成insert插入语句
	//u := User{ID: 1, Name: "Bob", Email: "bob@example.com"}
	//sql, vals := buildInsertSQL(u)
	//fmt.Println("生成 SQL:", sql)
	//fmt.Println("对应值:", vals)

	// 使用unsafe.pointer 去修改结构体的小写字段 ；实现零拷贝
}

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

//func buildInsertSQL(x interface{}) (string, []interface{}) {
//	val := reflect.ValueOf(x)
//
//}
