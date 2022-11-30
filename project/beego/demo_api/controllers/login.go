package controllers

import (
	"demo_api/models"
	"encoding/json"
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
