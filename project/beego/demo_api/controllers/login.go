package controllers

import (
	"demo_api/models"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type LoginController struct {
	beego.Controller
}

func (ctrl *LoginController) Resister() {
	user := models.Users{}
	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &user); err != nil {
		ctrl.Data["json"] = err.Error()
	}
	err := user.Insert(&user)
	if err != nil {
		return
	}
	ctrl.Data["json"] = user
	ctrl.ServeJSON()
}

func (ctrl *LoginController) Login() {
	req := make(map[string]string)
	data := ctrl.Ctx.Input.RequestBody //在RequestBody中读取Json
	json.Unmarshal(data, &req)

	logs.Info("%v", req)
	id, name := models.FindUser(req["email"], req["password"])
	if id > 0 {
		ctrl.Data["json"] = ReturnSuccess(0, "登录成功！", map[string]interface{}{"uid": id, "name": name}, 0)
	} else {
		ctrl.Data["json"] = ReturnError(7001, "登录失败！")
	}
	ctrl.ServeJSON()
}
