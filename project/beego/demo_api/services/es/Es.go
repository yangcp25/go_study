package es

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
func EsSearch(indexName string, query map[string]interface{}, from int, size int, sort map[string]string) HitsData {
	data := make(map[string]interface{})
	data["query"] = query
	data["from"] = from
	data["size"] = size
	data["sort"] = sort
	bytesData, _ := json.Marshal(data)
	url := esUrl + indexName + "/_search"
	resp, _ := http.Post(url, "application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)

	var stb ReqSearchData

	err := json.Unmarshal(body, &stb)

	if err != nil {
		fmt.Println(err)
	}
	return stb.hits
}

type ReqSearchData struct {
	hits HitsData `json:"hits"`
}

type HitsData struct {
	Total TotalData     `json:"total"`
	Hits  []HitsTwoData `json:"hits"`
}

type HitsTwoData struct {
	Source json.RawMessage `json:"_source"`
}

type TotalData struct {
	Value    int
	Relation string
}

func EsAdd(indexName string, id string, body map[string]interface{}) bool {
	url := esUrl + indexName + "/_dos/" + id
	_, err := HttpPost(url, body)
	if err != nil {
		return false
	}
	return true
}

func EsUpdate(indexName string, id string, body map[string]interface{}) bool {
	url := esUrl + indexName + "/_dos/" + id + "/_update"
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
	url := esUrl + indexName + "/_dos/" + id
	err := HttpDelete(url)
	if err != nil {
		return false
	}
	return true
}

func HttpPost(url string, body map[string]interface{}) ([]byte, error) {
	bytesData, _ := json.Marshal(body)
	resp, _ := http.Post(url, "application/json", bytes.NewReader(bytesData))
	res, err := ioutil.ReadAll(resp.Body)
	return res, err
}

func HttpDelete(url string) error {

	req, _ := http.NewRequest("DELETE", url, nil)

	_, err := http.DefaultClient.Do(req)

	return err
}
