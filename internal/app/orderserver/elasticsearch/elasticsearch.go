package elasticsearch

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"have-you-ordered/cmd/orderserver/app/config"
	"log"
)

var Client *elasticsearch.Client

var IndexName = config.Config.Index

type EsBaseResponseBody struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
}

func init() {
	cfg := elasticsearch.Config{
		Addresses: config.Config.Elasticsearch.Hosts,
		Username:  config.Config.Elasticsearch.Username,
		Password:  config.Config.Elasticsearch.Password,
	}
	Client, _ = elasticsearch.NewClient(cfg)

	var r map[string]interface{}
	res, err := Client.Info()
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
