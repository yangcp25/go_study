package main

import (
	"encoding/json"
	"fmt"
	"fromwork/web"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		str := []byte{'h', 'e', 'l', 'l', '0'}
		writer.Write(str)
	})

	http.HandleFunc("/test2", handler2)
	http.HandleFunc("/test3", handler3)
	http.HandleFunc("/returnJson", handler4)
	http.HandleFunc("/returnJson2", handler5)
	http.ListenAndServe("localhost:8081", nil)
}

func handler5(writer http.ResponseWriter, request *http.Request) {
	context := &web.Context{
		request, writer,
	}

	var data map[string]interface{}

	context.ReadJson(&data)

	context.WriteJson2(1, "成功", data)
}

func handler4(writer http.ResponseWriter, request *http.Request) {
	context := &web.Context{
		request, writer,
	}

	var data map[string]interface{}

	context.ReadJson(&data)

	//
	jsonData := map[string]interface{}{
		"code": 4,
		"msg":  "xxxxx",
		"data": data,
	}

	res, _ := json.Marshal(jsonData)

	_, err := writer.Write(res)
	if err != nil {
		fmt.Errorf("err" + err.Error())
	}
}

// 获取参数
func handler2(writer http.ResponseWriter, request *http.Request) {
	// body 获取
	/*body := request.Body
	data, _ := ioutil.ReadAll(body)

	var params map[string]interface{}
	err := json.Unmarshal(data, &params)
	if err != nil {
		fmt.Println("err:" + err.Error())
		return
	}
	fmt.Println(params)*/
	// form 获取 文件等
	//request.ParseMultipartForm(19999)
	//
	//fmt.Println(request.MultipartForm)
	//request.ParseForm()

	//fmt.Println(request.Form)
	//
	//// 获取query
	//data := request.URL.Query()
	//
	//fmt.Println(data["test"])
}

// 封装数据获取

func handler3(writer http.ResponseWriter, request *http.Request) {
	context := &web.Context{request, writer}

	var data map[string]interface{}
	context.ReadJson(&data)

	fmt.Println(data)
}
