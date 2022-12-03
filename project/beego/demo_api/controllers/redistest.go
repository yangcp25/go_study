package controllers

import (
	"demo_api/services"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
)

// Operations about Users
type RedisTestController struct {
	beego.Controller
}

func (r *RedisTestController) Test() {
	c := services.PoolConnect()
	defer c.Close()

	value, _ := redis.String(c.Do("get", "test"))

	if value == "" {
		_, err := c.Do("set", "test", "ycp")
		if err == nil {
			c.Do("expire", "test", 1000)
		}
	}

	ttl, _ := redis.Int64(c.Do("ttl", "test"))

	r.Data["json"] = ReturnSuccess(0, "", value, ttl)
	r.ServeJSON()
}
