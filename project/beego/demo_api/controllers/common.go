package controllers

import beego "github.com/beego/beego/v2/server/web"

type commonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) *JsonStruct {
	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Items: items,
		Count: count,
	}
	return json
}

func ReturnError(code int, msg interface{}) *JsonStruct {
	json := &JsonStruct{
		Code: code,
		Msg:  msg,
	}
	return json
}
