package es_package

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var esUrl string

func init() {
	esUrl = "http://192.168.1.3:19200/"
}
func EsSearch(indexName string, query map[string]interface{}, from int, size int, sort []map[string]string) map[string]interface{} {
	data := make(map[string]interface{})
	data["query"] = query
	data["from"] = from
	data["size"] = size
	//data["sort"] = sort
	bytesData, _ := json.Marshal(data)
	url := esUrl + indexName + "/_search"
	resp, _ := http.Post(url, "application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)

	var stb map[string]interface{}

	err := json.Unmarshal(body, &stb)

	if err != nil {
		fmt.Println(err)
	}
	return stb
}

type ReqSearchData struct {
	Hits HitsData `json:"hits.hits"`
}

type HitsData struct {
	Source []TotalData `json:"_source"`
}

type TotalData struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

func EsAdd(indexName string, id string, body map[string]interface{}) bool {
	url := esUrl + indexName + "/_doc/" + id
	res, err := HttpPost(url, body)
	if err != nil {
		return false
	}
	fmt.Println(url, res)
	return true
}

func EsUpdate(indexName string, id string, body map[string]interface{}) bool {
	url := esUrl + indexName + "/_doc/" + id + "/_update"
	updateData := map[string]interface{}{
		"doc": body,
	}
	_, err := HttpPost(url, updateData)
	if err != nil {
		return false
	}
	return true
}

func EsDelete(indexName string, id string) bool {
	url := esUrl + indexName + "/_doc/" + id
	err := HttpDelete(url)
	if err != nil {
		return false
	}
	return true
}

func HttpPost(url string, body map[string]interface{}) (string, error) {
	bytesData, _ := json.Marshal(body)
	resp, err := http.Post(url, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		fmt.Println(err)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(res), err
}

func HttpDelete(url string) error {

	req, _ := http.NewRequest("DELETE", url, nil)

	_, err := http.DefaultClient.Do(req)

	return err
}
