package es_package

import (
	"demo_api/models"
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

type EsSearchData struct {
	ShouldQuery []elastic.Query
	From        int
	Size        int
	Sorters     []elastic.Sorter
}

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
			"content":			{ "type": "text","analyzer": "ik_max_word"},
			"create_time":		{ "type": "text" }
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

func (es *UserES) Search(ctx context.Context, req map[string]string) ([]*models.Comments, error) {
	id := req["id"]
	create_time := req["create_time"]
	content := req["content"]

	var search EsSearchData

	if len(id) != 0 {
		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("id", content))
	}
	if len(create_time) != 0 {
		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("create_time", create_time))
	}
	if len(content) != 0 {
		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("content", content))
	}
	search.Sorters = append(search.Sorters, elastic.NewFieldSort("create_time").Desc())

	var size, _ = strconv.Atoi(req["size"])
	var num, _ = strconv.Atoi(req["num"])
	search.From = (num - 1) * size
	search.Size = size

	boolQuery := elastic.NewBoolQuery()
	boolQuery.Should(search.ShouldQuery...)

	service := es.client.Search().Index(es.index).Query(boolQuery).From(search.From).Size(search.Size)
	resp, err := service.Do(ctx)

	if err != nil {
		return nil, err
	}

	if resp.TotalHits() == 0 {
		return nil, nil
	}

	fmt.Println(resp.TotalHits())
	userES := make([]*models.Comments, 0)
	for _, e := range resp.Each(reflect.TypeOf(&models.Comments{})) {
		us := e.(*models.Comments)
		userES = append(userES, us)
	}

	return userES, nil
}
