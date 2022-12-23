package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type DemoController struct {
	beego.Controller
}

func (this *DemoController) Hello() {
	str := "hello world!"
	this.Ctx.WriteString(str)
}
