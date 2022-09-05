package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:12306)/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)

	_, err := rand.Read(u[:])

	if err != nil {
		panic(err)
	}

	u[8] = (u[8] | 0x40) & 0x7F

	u[6] = (u[6] & 0xF) | (0x4 << 4)

	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// 哈希加密
func Encrypt(plainText string) (cryptText string) {
	cryptText = fmt.Sprintf("%x", sha1.Sum([]byte(plainText)))
	return
}
