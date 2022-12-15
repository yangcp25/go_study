package controllers

import (
	"bytes"
	"context"
	"demo_api/services/es_package"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strconv"
	"strings"
	"sync"
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
					"_id": "index_1",
				},
			},
		},
	}

	res := es_package.EsSearch("test_index_1", query, 0, 10, sort)

	fmt.Println(res)
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *EsDemoController) EsSearch2() {

	address := []string{"http://192.168.1.3:19200"}

	config := elasticsearch.Config{
		Addresses: address,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, _ := elasticsearch.NewClient(config)

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "测试title",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Println("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test_index_2"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		fmt.Println("Error encoding query2: %s", err)
	}

	fmt.Println(res)
	defer res.Body.Close()
	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	this.Data["json"] = r
	this.ServeJSON()
}

type EsData struct {
	Source SourceData `json:"_source"`
}

type SourceData struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

func (c EsDemoController) test2() {
	log.SetFlags(0)

	var (
		r  map[string]interface{}
		wg sync.WaitGroup
	)

	// Initialize a client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//
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
		log.Fatalf("Error creating the client: %s", err)
	}

	// 1. Get cluster info
	//
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	// 2. Index documents concurrently
	//
	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			data, err := json.Marshal(struct{ Title string }{Title: title})
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       bytes.NewReader(data),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "测试title",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}
