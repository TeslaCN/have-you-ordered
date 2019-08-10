package elasticsearch

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var Es *elasticsearch.Client

const IndexName = "having-meal"

type EsBaseResponseBody struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	//Hits struct {
	//	Total struct {
	//		Value    uint64 `json:"value"`
	//		Relation string `json:"relation"`
	//	}
	//	MaxScore uint64     `json:"max_score"`
	//	Hits     []struct{} `json:"hits"`
	//} `json:"hits"`
	//Aggregations interface{} `json:"aggregations"`
}

func init() {
	// TODO 抽取到配置项
	config := elasticsearch.Config{
		Addresses: []string{"http://es.0:49204"},
	}
	Es, _ = elasticsearch.NewClient(config)

	var r map[string]interface{}
	res, err := Es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
}
