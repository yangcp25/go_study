package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Context struct {
	R *http.Request
	W http.ResponseWriter
}

func (c *Context) ReadJson(data interface{}) {
	body := c.R.Body

	res, err := ioutil.ReadAll(body)

	if err != nil {
		_ = fmt.Errorf("读取数据错误！")
	}

	err = json.Unmarshal(res, data)
	if err != nil {
		_ = fmt.Errorf("序列化失败！")
	}
}

func (c Context) WriteJson(status int8, msg string, data interface{}) error {
	//
	//jsonData := msgInfo{
	//	code: status,
	//	msg:  msg,
	//	data: data,
	//}
	jsonData := map[string]interface{}{
		"code": status,
		"msg":  msg,
		"data": data,
	}
	res, err := json.Marshal(jsonData)

	if err != nil {
		return fmt.Errorf("序列化失败")
	}
	_, err = c.W.Write(res)
	if err != nil {
		return err
	}

	return nil
}

type msgInfo struct {
	Code int
	Msg  string
	Data interface{}
}

func (c Context) WriteJson2(status int, msg string, data interface{}) error {
	//
	jsonData := &msgInfo{
		Code: status,
		Msg:  msg,
		Data: data,
	}

	res, err := json.Marshal(jsonData)

	if err != nil {
		return fmt.Errorf("序列化失败")
	}

	_, err = c.W.Write(res)

	c.W.WriteHeader(http.StatusOK)
	c.W.Header().Set("content-type", "application/json")
	if err != nil {
		return err
	}

	return nil
}
