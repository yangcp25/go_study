package services

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func Connect() redis.Conn {
	pool, _ := redis.Dial("tcp", "127.0.0.1:16379")
	return pool
}

func PoolConnect() redis.Conn {
	pool := &redis.Pool{
		MaxIdle:     1,  //最大的空闲连接数
		MaxActive:   10, //最大连接数
		IdleTimeout: 180 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, _ := redis.Dial("tcp", "127.0.0.1:16379")
			return c, nil
		},
	}
	return pool.Get()
}
