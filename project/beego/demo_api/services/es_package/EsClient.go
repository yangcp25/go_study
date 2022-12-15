package es_package

import (
	"demo_api/models"
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
	"time"
)

func NewEsClient() *elastic.Client {
	url := fmt.Sprintf("http://%s:%d", "192.168.1.3", 9200)
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetSniff(false),
		elastic.SetURL(url),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		log.Fatalf("Failed to create elastic client, url%v\n", url)
	}
	return client
}

var mappingTpl = `{
	"mappings":{
		"properties":{
			"id": 				{ "type": "long" },
			"userid": 		{ "type": "long" },
			"content":			{ "type": "text" },
			"create_time":		{ "type": "text" },
			}
		}
	}`

type UserES struct {
	client  *elastic.Client
	index   string
	mapping string
}

func CreateIndex(client *elastic.Client, indexName string) *UserES {
	userEs := &UserES{
		client:  client,
		index:   indexName,
		mapping: mappingTpl,
	}

	userEs.init()

	return userEs
}

func (es *UserES) init() {
	ctx := context.Background()

	exists, err := es.client.IndexExists(es.index).Do(ctx)
	if err != nil {
		fmt.Printf("userEs init exist failed err is %s\n", err)
		return
	}

	if !exists {
		_, err := es.client.CreateIndex(es.index).Body(es.mapping).Do(ctx)
		if err != nil {
			fmt.Printf("userEs init failed err is %s\n", err)
			return
		}
	}
}

func (es *UserES) BatchAdd(ctx context.Context, comments []*models.Comments) error {
	var err error
	for i := 0; i < 3; i++ {
		if err = es.batchAdd(ctx, comments); err != nil {
			fmt.Println("batch add failed ", err)
			continue
		}
		return err
	}
	return err
}

func (es *UserES) batchAdd(ctx context.Context, comments []*models.Comments) error {
	req := es.client.Bulk().Index(es.index)
	for _, u := range comments {
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)
		u.CreateTime = tm.Format("2006-01-02 03:04:05")
		doc := elastic.NewBulkIndexRequest().Id(strconv.FormatUint(uint64(u.Id), 10)).Doc(u)
		req.Add(doc)
	}
	if req.NumberOfActions() < 0 {
		return nil
	}
	if _, err := req.Do(ctx); err != nil {
		return err
	}
	return nil
}
