package controllers

import (
	"demo_api/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
	"net/http"
)

// Operations about Users
type CommentsController struct {
	beego.Controller
}

func (this *CommentsController) SaveComments() {

	var req models.Comments
	_ = json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	userId := req.UserId
	content := req.Content

	id, err := models.SaveComments(content, userId)

	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "保存成功", id, 0)
	} else {
		this.Data["json"] = ReturnError(5000, "保存失败")
	}

	this.ServeJSON()
}

func (this *CommentsController) UpdateComments() {
	var req models.Comments
	_ = json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	id := req.Id
	content := req.Content

	id, err := models.EditComments(id, content)

	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "修改成功", id, 0)
	} else {
		this.Data["json"] = ReturnError(5000, "修改失败")
	}

	this.ServeJSON()
}

func (this *CommentsController) DeleteComments() {
	var req models.Comments
	_ = json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	content := req.Content

	res, err := models.DeleteComments(content)

	if err == nil && res > 0 {
		this.Data["json"] = ReturnSuccess(0, "删除成功", res, 0)
	} else {
		this.Data["json"] = ReturnError(5000, "删除失败")
	}

	this.ServeJSON()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *CommentsController) BuildSockets() {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	if conn, err = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil); err != nil {
		goto ERR
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
