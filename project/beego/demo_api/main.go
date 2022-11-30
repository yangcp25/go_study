package main

import (
	_ "demo_api/routers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}
	databaseConfig, _ := config.String("sqlconn")
	logs.Info("数据库配置", databaseConfig)
	err := orm.RegisterDataBase("default", "mysql", databaseConfig)
	if err != nil {
		logs.Info("数据库初始化错误", err)
		return
	}
	beego.Run()
}
