// exercise10.go
package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func buildInsertSQL(x interface{}) (string, []interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Struct {
		return "", nil
	}
	t := v.Type()
	var cols []string
	var placeholders []string
	var values []interface{}
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			continue
		}
		cols = append(cols, tag)
		placeholders = append(placeholders, "?")
		values = append(values, v.Field(i).Interface())
	}
	tableName := strings.ToLower(t.Name()) // 简单用 struct 名称小写做表名
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		tableName,
		strings.Join(cols, ", "),
		strings.Join(placeholders, ", "),
	)
	return sql, values
}

func main() {
	u := User{ID: 1, Name: "Bob", Email: "bob@example.com"}
	sql, vals := buildInsertSQL(u)
	fmt.Println("生成 SQL:", sql)
	fmt.Println("对应值:", vals)
}
