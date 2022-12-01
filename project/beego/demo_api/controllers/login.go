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

	id, name := models.FindUser(req["email"], req["password"])
	if id > 0 {
		ctrl.Data["json"] = ReturnSuccess(0, "登录成功！", map[string]interface{}{"uid": id, "name": name}, 0)
	} else {
		ctrl.Data["json"] = ReturnError(7001, "登录失败！")
	}
	ctrl.ServeJSON()
}

func (this *LoginController) GetUserList() {

	req := make(map[string]int)
	data := this.Ctx.Input.RequestBody //在RequestBody中读取Json
	json.Unmarshal(data, &req)

	logs.Info("%v", req)
	//limit, _ := strconv.Atoi(req["limit"])
	//offset, _ := strconv.Atoi(req["offset"])
	//remember_token, _ := strconv.Atoi(req["remember_token"])
	limit, _ := req["limit"]
	offset, _ := req["offset"]
	remember_token, _ := req["remember_token"]
	if limit == 0 {
		limit = 2
	}

	count, users, err := models.GetUserList(limit, offset, remember_token)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "获取成功！", users, count)
	} else {
		this.Data["json"] = ReturnError(7001, "获取失败！")
	}
	this.ServeJSON()
}
