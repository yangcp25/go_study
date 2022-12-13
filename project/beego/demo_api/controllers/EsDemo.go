package controllers

import (
	"bytes"
	"demo_api/services/es_package"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/elastic/go-elasticsearch/v7"
)

type EsDemoController struct {
	beego.Controller
}

func (c *EsDemoController) AddIndex() {
	address := []string{"http://192.168.1.3:19200"}

	config := elasticsearch.Config{
		Addresses: address,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)

	if err != nil {
		fmt.Println(err)
	}

	doc := map[string]interface{}{
		"title":   "测试title",
		"content": "测试一下啊",
	}

	byteData, err := json.Marshal(doc)

	if err != nil {
		fmt.Println(err)
	}

	index, err := es.Index("test_index_1", bytes.NewReader(byteData))
	if err != nil {
		fmt.Println(err)
	}

	defer index.Body.Close()

	c.Data["str"] = index.String()
	c.ServeJSON()
}

func (this *EsDemoController) Add() {
	body := map[string]interface{}{
		"title":   "heihie",
		"content": "hamapo",
	}

	res := es_package.EsAdd("test_index_1", "index_1", body)

	this.Data["res"] = res
	this.ServeJSON()
}

func (this *EsDemoController) Edit() {
	body := map[string]interface{}{
		"title":   "heihie2",
		"content": "hamapo2",
	}

	res := es_package.EsUpdate("test_index_1", "index_1", body)

	this.Data["res"] = res
	this.ServeJSON()
}

func (this *EsDemoController) EsDelete() {

	res := es_package.EsDelete("test_index_1", "JsCxC4UBAWFfYdkanMJ2")

	this.Data["res"] = res
	this.ServeJSON()
}

func (this *EsDemoController) EsSearch() {

	sort := []map[string]string{{
		"id": "desc",
	}}

	query := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"term": map[string]interface{}{
					"id": "index_1",
				},
			},
		},
	}

	res := es_package.EsSearch("test_index_1", query, 0, 10, sort)

	fmt.Println(res)
	this.Data["res"] = res
	this.ServeJSON()
}
